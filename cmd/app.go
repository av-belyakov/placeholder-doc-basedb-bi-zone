package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"

	"golang.org/x/sync/errgroup"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/constants"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/interfaces"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/dicontainer"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/supportingfunctions"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/wrappers"
)

type App struct {
	diContainer *dicontainer.DiContainer
	appRouter   *ApplicationRouter
	ctx         context.Context
}

func NewApp(ctx context.Context) *App {
	rootPath, err := supportingfunctions.GetPathRoot(constants.Root_Dir)
	if err != nil {
		log.Fatalf("error, it is impossible to form root path (%s)", err.Error())
	}

	ch := make(chan interfaces.Messager)
	app := &App{
		ctx:         ctx,
		diContainer: dicontainer.NewDIContainer(rootPath, ch),
	}
	app.appRouter = NewRouter(
		app.diContainer.Logger(ctx),
		app.diContainer.Counter(ctx),
		ApplicationRouterSettings{
			ChanToNats:    app.diContainer.NatsConnecter(ctx).GetChannelToModule(),
			ChanFromNats:  app.diContainer.NatsConnecter(ctx).GetChannelFromModule(),
			ChanToKafka:   app.diContainer.KafkaConnecter(ctx).GetChannelToModule(),
			ChanFromKafka: app.diContainer.KafkaConnecter(ctx).GetChannelFromModule(),
			ChanToDBS:     app.diContainer.DB(ctx).GetChannelToModule(),
			ChanFromDBS:   app.diContainer.DB(ctx).GetChannelFromModule(),
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
	wrappers.WrappersZabbixInteraction(ctx, zabbixSettings, app.diContainer.SimpleLogger(ctx), ch)

	return app
}

func (a *App) Start() {
	// сервер отладки
	if a.diContainer.Configer().GetDebugServer().Enable {
		go func() {
			httpServer := &http.Server{
				Addr: fmt.Sprintf(
					"%s:%d",
					a.diContainer.Configer().GetDebugServer().Host,
					a.diContainer.Configer().GetDebugServer().Port,
				),
				BaseContext: func(_ net.Listener) context.Context {
					return a.ctx
				},
			}

			g, gCtx := errgroup.WithContext(a.ctx)
			g.Go(func() error {
				return httpServer.ListenAndServe()
			})
			g.Go(func() error {
				<-gCtx.Done()

				return httpServer.Shutdown(context.Background())
			})

			if err := g.Wait(); err != nil {
				log.Fatal("error debugging server:", err)
			}
		}()
	}

	// вывод информационного сообщения при старте приложения
	msg := getInformationMessage(a.diContainer.Configer().Get())
	a.diContainer.SimpleLogger(a.ctx).Write("info", strings.ToLower(msg))

	a.appRouter.Router(a.ctx)

	<-a.ctx.Done()
}
