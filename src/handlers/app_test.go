package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/huyrun/counting-mouse/src/config"
	"github.com/stretchr/testify/require"
	"testing"
)

func newTestApp(t *testing.T, ctx context.Context) (*App, error) {
	config, err := config.NewAppConfig()
	require.NoError(t, err)

	app := &App{
		Config: config,
		Engine: gin.New(),
	}

	return app, app.setup(ctx)
}
