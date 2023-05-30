//go:build wireinject
// +build wireinject

package setup

import (
	"context"

	"github.com/google/wire"
	"github.com/huyrun/counting-mouse/src/handlers"
)

func NewApp(ctx context.Context) (*handlers.App, func(), error) {
	wire.Build(AppSet, handlers.App{})
	return nil, nil, nil
}

var AppSet = wire.NewSet(
	ProvideAppConfig,
	ProvideUseCase,
	ProvideEngine,
)
