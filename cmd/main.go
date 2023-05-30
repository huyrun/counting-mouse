package main

import (
	"context"
	"errors"
	"github.com/huyrun/counting-mouse/src/setup"
	"net/http"
)

func main() {
	ctx := context.Background()
	app, cleanup, err := setup.NewApp(ctx)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	if err = app.Start(ctx); err != nil && !errors.Is(err, http.ErrServerClosed) {
		panic(err)
	}
}
