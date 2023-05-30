package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/huyrun/counting-mouse/src/model"
	"github.com/huyrun/counting-mouse/src/util"
	"net/http"
)

// handlePing implements ping request
func (app *App) handlePing(ctx *gin.Context) {
	util.SendJSON(ctx, http.StatusOK, model.VerdictSuccess, "pong", map[string]interface{}{
		"service": "counting-mouse",
	})
}
