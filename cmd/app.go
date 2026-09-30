package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	_ "net/http/pprof"
	"strings"

	"golang.org/x/sync/errgroup"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/constants"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/interfaces"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/dicontainer"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/supportingfunctions"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/wrappers"
)

type App struct {
	chanMessage chan interfaces.Messager
	diContainer *dicontainer.DiContainer
	appRouter   *ApplicationRouter
}

// NewApp конструктор
func NewApp() *App {
	rootDirPath, err := supportingfunctions.GetPathRoot(constants.Root_Dir)
	if err != nil {
		log.Fatalf("error, it is impossible to form root dir path (%s)", err.Error())
	}

	ch := make(chan interfaces.Messager)
	app := &App{
		chanMessage: ch,
		diContainer: dicontainer.NewDIContainer(rootDirPath, ch),
	}

	return app
}

// Start инициализация запуска приложения
func (app *App) Start(ctx context.Context) {
	// инициализация внутреннего роутера
	app.appRouter = NewRouter(
		app.diContainer.Logger(ctx),
		app.diContainer.Counter(ctx),
		ApplicationRouterSettings{
			ChanToNats:    app.diContainer.NatsConnecter(ctx).GetChannelToModule(),
			ChanFromNats:  app.diContainer.NatsConnecter(ctx).GetChannelFromModule(),
			ChanToKafka:   app.diContainer.KafkaConnecter(ctx).GetChannelToModule(),
			ChanFromKafka: app.diContainer.KafkaConnecter(ctx).GetChannelFromModule(),
			ChanToDBS:     app.diContainer.DbConnecter(ctx).GetChannelToModule(),
			ChanFromDBS:   app.diContainer.DbConnecter(ctx).GetChannelFromModule(),
		})

	// настройка обёртки для взаимодействия с Zabbix
	zabbixSettings := wrappers.WrappersZabbixInteractionSettings{
		NetworkPort: app.diContainer.Configer().GetCommon().Zabbix.NetworkPort,
		NetworkHost: app.diContainer.Configer().GetCommon().Zabbix.NetworkHost,
		ZabbixHost:  app.diContainer.Configer().GetCommon().Zabbix.ZabbixHost,
		EventTypes:  make([]wrappers.EventType, len(app.diContainer.Configer().GetCommon().Zabbix.EventTypes)),
	}
	for _, v := range app.diContainer.Configer().GetCommon().Zabbix.EventTypes {
		zabbixSettings.EventTypes = append(zabbixSettings.EventTypes, wrappers.EventType{
			IsTransmit: v.IsTransmit,
			EventType:  v.EventType,
			ZabbixKey:  v.ZabbixKey,
			Handshake: wrappers.Handshake{
				TimeInterval: v.Handshake.TimeInterval,
				Message:      v.Handshake.Message,
			},
		})
	}
	// обертка для взаимодействия с Zabbix
	wrappers.WrappersZabbixInteraction(ctx, zabbixSettings, app.diContainer.SimpleLogger(ctx), app.chanMessage)

	// запуск сервера отладки
	if app.diContainer.Configer().GetDebugServer().Enable {
		go app.startDebugServer(ctx)
	}

	// вывод информационного сообщения при старте приложения
	msg := getInformationMessage(app.diContainer.Configer().Get())
	app.diContainer.SimpleLogger(ctx).Write("info", strings.ToLower(msg))

	app.appRouter.Start(ctx)

	<-ctx.Done()
}

func (app *App) startDebugServer(ctx context.Context) {
	httpServer := &http.Server{
		Addr: fmt.Sprintf(
			"%s:%d",
			app.diContainer.Configer().GetDebugServer().Host,
			app.diContainer.Configer().GetDebugServer().Port,
		),
		BaseContext: func(_ net.Listener) context.Context {
			return ctx
		},
	}

	g, gCtx := errgroup.WithContext(ctx)
	g.Go(func() error {
		return httpServer.ListenAndServe()
	})
	g.Go(func() error {
		<-gCtx.Done()

		return httpServer.Shutdown(context.Background())
	})

	app.appRouter.logger.Send(
		"info",
		fmt.Sprintf(
			"start debug server with address '%s' port '%d'",
			app.diContainer.Configer().GetDebugServer().Host,
			app.diContainer.Configer().GetDebugServer().Port,
		))

	if err := g.Wait(); err != nil {
		app.appRouter.logger.Send("error", fmt.Sprintf("error debugging server: '%s'", err.Error()))
	}
}
