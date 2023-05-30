package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/huyrun/counting-mouse/src/model"
	"github.com/huyrun/counting-mouse/src/util"
	"net/http"
)

type countAliveMousesRequest struct {
	OriginalAmount int `json:"originalAmount" form:"originalAmount" binding:"required"`
	Months         int `json:"months" form:"months" binding:"required"`
	LifeSpan       int `json:"lifeSpan" form:"lifeSpan" binding:"required"`
}

func (app *App) handleCountingAliveMouses(ctx *gin.Context) {
	var req countAliveMousesRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		util.SendJSON(ctx, http.StatusBadRequest, model.VerdictInvalidParameters, "invalid parameters", nil)
		return
	}

	aliveMouseAmount := app.Usecase.CountAliveMouse(ctx, req.OriginalAmount, req.LifeSpan, req.Months)

	util.SendJSON(ctx, http.StatusOK, model.VerdictSuccess, "success", map[string]interface{}{
		"alive_mouse_amount": aliveMouseAmount,
	})
}
