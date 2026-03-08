package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func addUserRouter(rg *gin.RouterGroup) {
	users := rg.Group("/users")

	users.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "users")
	})

	users.GET("/comments", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "users comments")
	})
	users.GET("/pictures", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "users pictures")
	})
}
