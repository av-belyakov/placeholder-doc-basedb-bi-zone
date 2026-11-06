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

	// запуск сервера отладки
	if app.diContainer.Configer().GetDebugServer().Enable {
		go app.startDebugServer(ctx)
	}

	// вывод информационного сообщения при старте приложения
	msg := getInformationMessage(app.diContainer.Configer().Get())
	app.diContainer.Logger(ctx).Send("info", strings.ToLower(msg))

	if err := app.appRouter.Start(ctx); err != nil {
		log.Fatalf("launching the application router %s", err.Error())
	}

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
		log.Printf("%vthe debugging server is stopped%v\n", constants.Ansi_Dark_Red, constants.Ansi_Reset)

		return httpServer.Shutdown(context.Background())
	})

	app.diContainer.Logger(ctx).Send(
		"info",
		fmt.Sprintf(
			"start debug server with address '%s' port '%d'",
			app.diContainer.Configer().GetDebugServer().Host,
			app.diContainer.Configer().GetDebugServer().Port,
		))

	if err := g.Wait(); err != nil {
		app.diContainer.Logger(ctx).Send("error", fmt.Sprintf("start debug server - '%s'", err.Error()))
	}
}
