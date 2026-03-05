package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "pong")
	})

	r.GET("/user/:name", func(ctx *gin.Context) {
		user := ctx.Params.ByName("name")
		value, ok := db[user]
		if ok {
			ctx.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		}
	})

	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar",
		"manu": "123",
	}))

	authorized.POST("admin", func(ctx *gin.Context) {
		// 从上下文获取通过认证的用户名（gin.AuthUserKey 是内置的键常量）
		user := ctx.MustGet(gin.AuthUserKey).(string)

		// 定义临时结构体，用于绑定客户端传入的 JSON 数据
		var json struct {
			Value string `json:"value" binding:"required"` // 必传字段校验
		}
		// 绑定 JSON 数据到结构体（自动校验 "required" 规则）
		if ctx.Bind(&json) == nil {
			db[user] = json.Value // 将值存入内存 map
			ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	})

	return r
}

func main() {
	r := setupRouter()
	_ = r.Run(":8080")
}
