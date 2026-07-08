package handlers

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping(ctx *gin.Context) {
	slog.Debug("ping endpoint is called")
	ctx.JSON(http.StatusOK, map[string]any{"ping": "pong"})

}

func Health(ctx *gin.Context) {
	slog.Debug("health endpoint is called")
	ctx.JSON(http.StatusOK, gin.H{
		"health": "ok",
	})
}
