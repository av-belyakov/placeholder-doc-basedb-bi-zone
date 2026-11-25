package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/pprof"
	_ "net/http/pprof"
	"strings"

	"golang.org/x/sync/errgroup"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/constants"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/dicontainer"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/supportingfunctions"
)

type App struct {
	diContainer *dicontainer.DiContainer
	appRouter   *ApplicationRouter
}

// NewApp конструктор
func NewApp() *App {
	rootDirPath, err := supportingfunctions.GetPathRoot(constants.Root_Dir)
	if err != nil {
		log.Fatalf("getting the root directory %s", err.Error())
	}

	app := &App{
		diContainer: dicontainer.NewDIContainer(rootDirPath),
	}

	return app
}

// Start инициализация запуска приложения
func (app *App) Start(ctx context.Context) {
	// инициализация внутреннего роутера
	app.appRouter = NewRouter(app.diContainer)

	// запуск сервера сервисного монитора
	if app.diContainer.Configer().GetServiceMonitor().Enable {
		go app.startServiceMonitor(ctx)
	}

	// вывод информационного сообщения при старте приложения
	msg := getInformationMessage(app.diContainer.Configer().Get())
	app.diContainer.Logger(ctx).Send("info", strings.ToLower(msg))

	if err := app.appRouter.Start(ctx); err != nil {
		log.Fatalf("launching the application router %s", err.Error())
	}

	<-ctx.Done()
}

func (app *App) startServiceMonitor(ctx context.Context) {
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.HandlerFor(app.diContainer.PrometheusConnecter().GetRegistry(), promhttp.HandlerOpts{
		EnableOpenMetrics: true,
	}))
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/heap", pprof.Handler("heap").ServeHTTP)
	mux.HandleFunc("/debug/pprof/goroutine", pprof.Handler("goroutine").ServeHTTP)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)

	server := &http.Server{
		Addr: fmt.Sprintf(
			"%s:%d",
			app.diContainer.Configer().GetServiceMonitor().Host,
			app.diContainer.Configer().GetServiceMonitor().Port,
		),
		BaseContext: func(_ net.Listener) context.Context {
			return ctx
		},
		Handler: mux,
	}

	g, gCtx := errgroup.WithContext(ctx)
	g.Go(func() error {
		return server.ListenAndServe()
	})
	g.Go(func() error {
		<-gCtx.Done()
		log.Printf("%vthe service monitor server server is stopped%v\n", constants.Ansi_Dark_Red, constants.Ansi_Reset)

		return server.Shutdown(context.Background())
	})

	app.diContainer.Logger(ctx).Send(
		"info",
		fmt.Sprintf(
			"start service monitor server with address '%s' port '%d'",
			app.diContainer.Configer().GetServiceMonitor().Host,
			app.diContainer.Configer().GetServiceMonitor().Port,
		))

	if err := g.Wait(); err != nil {
		app.diContainer.Logger(ctx).Send("error", fmt.Sprintf("start service monitor server - '%s'", err.Error()))
	}
}
