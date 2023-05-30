package handlers

import "net/http"

func (app *App) setupRoutes() {
	en := app.Engine
	en.Handle(http.MethodGet, "/ping", app.handlePing)

	v1 := en.Group("/v1")
	{
		v1.Handle(http.MethodGet, "/mouses/alive", app.handleCountingAliveMouses)
	}
}
