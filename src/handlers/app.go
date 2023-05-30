package handlers

import (
	"context"
	"github.com/huyrun/counting-mouse/src/config"
	"github.com/huyrun/counting-mouse/src/usecase"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

type AppOption func(*App)

type App struct {
	Config  *config.AppConfig
	Engine  *gin.Engine
	Usecase usecase.UseCase
}

func (app *App) Start(ctx context.Context) error {
	if err := app.setup(ctx); err != nil {
		return err
	}

	ctx, cancel := signal.NotifyContext(ctx, syscall.SIGTERM, syscall.SIGKILL)
	defer cancel()

	srv := &http.Server{
		Addr:    app.Config.ServerAddress,
		Handler: app.Engine,
	}

	go func() {
		<-ctx.Done()

		shutdownCtx, done := context.WithTimeout(context.Background(), 5*time.Second)
		defer done()

		_ = srv.Shutdown(shutdownCtx)
	}()

	return srv.ListenAndServe()
}

func (app *App) setup(ctx context.Context) error {
	app.setupMiddlewares()
	app.setupRoutes()
	return nil
}

func (app *App) setupMiddlewares() {
}
