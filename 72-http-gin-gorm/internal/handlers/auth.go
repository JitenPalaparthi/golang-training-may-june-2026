package handlers

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Auth(ctx *gin.Context) {
	//if ctx.Request.Method != http.MethodGet {
	username := ctx.GetHeader("username")
	password := ctx.GetHeader("password")

	if username == "admin" && password == "admin123" {
		ctx.Next()
	} else {
		slog.Warn("unauthorized resource access")
		ctx.String(http.StatusUnauthorized, "unauthorized access")
		ctx.Abort()
	}
	//}
	//ctx.Next()
}
