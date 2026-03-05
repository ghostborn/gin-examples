package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func CookieTool() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		cookieValue, err := ctx.Cookie("label")
		if err != nil {
			ctx.JSON(http.StatusForbidden, gin.H{"error": "Forbidden with no cookie"})
			ctx.Abort()
			return
		}

		// 解析 Cookie 值（格式：ok_时间戳，比如 ok_1741234567）
		parts := strings.Split(cookieValue, "_")
		if len(parts) != 2 || parts[0] != "ok" {
			ctx.JSON(http.StatusForbidden, gin.H{"error": "Invalid cookie value"})
			ctx.Abort()
			return
		}

		// 解析时间戳并校验是否过期（30秒）
		timestamp, err := strconv.ParseInt(parts[1], 10, 64)
		if err != nil || time.Now().Unix()-timestamp > 30 {
			ctx.JSON(http.StatusForbidden, gin.H{"error": "Cookie expired"})
			ctx.Abort()
			return
		}

		// 校验通过
		ctx.Next()


	}
}

func main() {
	route := gin.Default()

	route.GET("/login", func(ctx *gin.Context) {
		// Cookie 值：ok + 当前时间戳
		cookieValue := fmt.Sprintf("ok_%d", time.Now().Unix())

		ctx.SetCookie("label", cookieValue, 30, "/", "localhost", false, true)
		ctx.String(200, "Login success!")
	})

	route.GET("/home", CookieTool(), func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"data": "Your home page"})
	})

	_ = route.Run(":8080")
}
