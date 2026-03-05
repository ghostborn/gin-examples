package main

import (
	"embed"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

//go:embed assets/* templates/*
var f embed.FS

func main() {
	router := gin.Default()
	templ := template.Must(template.New("").ParseFS(f, "templates/*.tmpl", "templates/foo/*.tmpl"))
	router.SetHTMLTemplate(templ)

	router.StaticFS("/public", http.FS(f))

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})

	router.GET("/foo", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "foo website",
		})
	})

	router.GET("favicon.ico", func(ctx *gin.Context) {
		file, _ := f.ReadFile("assets/favicon.ico")
		ctx.Data(
			http.StatusOK,
			"image/x-icon",
			file,
		)
	})

	_ = router.Run(":8080")

}
