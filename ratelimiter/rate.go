package main

import (
	"flag"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/ratelimit"
)

var (
	limit ratelimit.Limiter
	rps   = flag.Int("rps", 100, "request per second")
)

func setupLogging() {
	log.SetFlags(0)                  // 清空日志默认格式（如时间、文件名），只保留自定义前缀
	log.SetPrefix("[GIN] ")          // 日志前缀，和 Gin 默认日志风格一致
	log.SetOutput(gin.DefaultWriter) // 日志输出到 Gin 默认的 Writer（终端）
}

func leakBucket() gin.HandlerFunc {
	prev := time.Now	// 记录上一次请求的时间，初始为程序启动时
	return func(ctx *gin.Context) {
		now := limit.Take()
		
	}
}
