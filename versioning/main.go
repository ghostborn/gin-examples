package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	apiV1 := router.Group("/v1")
	apiV1.GET("users", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "List Of V1 Users")
	})

	authV1 := apiV1.Group("/", AuthMiddleWare())
	authV1.POST("users/add", AddV1User)

	apiV2 := router.Group("/v2")
	apiV2.GET("users", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "List Of V2 Users")
	})

	authV2 := apiV2.Group("/", AuthMiddleWare())
	authV2.POST("users/add", AddV2User)

	_ = router.Run(":8081")

}

func AddV1User(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "V1 User added")
}

func AddV2User(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "V2 User added")
}

func AuthMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		username := ctx.PostForm("user")
		password := ctx.PostForm("password")

		if username == "foo" && password == "bar" {
			return
		} else {
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
