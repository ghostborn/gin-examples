package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func addPingRoutes(rg *gin.RouterGroup) {
	ping := rg.Group("/ping")

	ping.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "pong")
	})
}
