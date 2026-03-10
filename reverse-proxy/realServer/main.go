package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

const (
	Addr = "127.0.0.1:2003"
)

func main() {
	r := gin.Default()
	r.GET("/:path", func(ctx *gin.Context) {
		req := ctx.Request
		urlPath := fmt.Sprintf("http://%s%s", Addr, req.URL.Path)
		realIP := fmt.Sprintf("RemoteAddr=%s,X-Forwarded-For=%v,X-Real-Ip=%v",
			req.RemoteAddr,
			req.Header.Get("X-Forwarded-For"),
			req.Header.Get("X-Real-Ip"),
		)
		ctx.JSON(200, gin.H{
			"path": urlPath,
			"ip":   realIP,
		})
	})

	if err := r.Run(Addr); err != nil {
		log.Printf("Error: %v", err)
	}
}
