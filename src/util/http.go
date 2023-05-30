package util

import (
	"github.com/gin-gonic/gin"
	"time"
)

// SendJSON sends JSON
func SendJSON(ctx *gin.Context, statusCode int, verdict string, message string, data interface{}) {
	ctx.Set("status_code", statusCode)
	ctx.Set("verdict", verdict)
	ctx.Set("message", message)
	ctx.Set("data", data)

	ctx.JSON(statusCode, gin.H{
		"verdict": verdict,
		"message": message,
		"data":    data,
		"time":    time.Now().Format(time.RFC3339),
	})
}
