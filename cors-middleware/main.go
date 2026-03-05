package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().
			Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		ctx.Writer.Header().
			Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	}
}

func main() {
	r := gin.Default()
	r.Use(CORSMiddleware())
	r.GET("/ping", func(ctx *gin.Context) {
		log.Println("Received a GET / ping request.")
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
			"time":    time.Now().Format(time.RFC3339),
		})
	})

	r.POST("/data", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"action":  "Created",
			"message": "Data processed successfully via POST.",
		})
	})

	r.PUT("/data/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		ctx.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"action":  "Updated",
			"message": fmt.Sprintf("Successfully processed PUT for resource ID: %s", id),
		})
	})

	r.DELETE("/data/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"action":  "Deleted",
			"message": fmt.Sprintf("Successfully processed DELETE for resource ID: %s", id),
		})
	})

	// 启动服务，监听 8080 端口
	// log.Fatal：启动失败时打印错误并退出程序（如端口被占用）
	log.Fatal(r.Run(":8080"))
}
