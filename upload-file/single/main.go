package main

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.MaxMultipartMemory = 8 << 20
	router.Static("/", "./public")
	router.POST("/upload", func(ctx *gin.Context) {
		name := ctx.PostForm("name")
		email := ctx.PostForm("email")

		file, err := ctx.FormFile("file")
		if err != nil {
			ctx.String(http.StatusBadRequest, "get form err: %s", err.Error())
			return
		}

		filename := filepath.Base(file.Filename)
		if err := ctx.SaveUploadedFile(file, filename); err != nil {
			ctx.String(http.StatusBadRequest, "upload file err: %s", err.Error())
			return
		}

		ctx.String(
			http.StatusOK,
			"File %s uploaded successfully with fields name=%s and email=%s.",
			file.Filename,
			name,
			email,
		)

	})
	_ = router.Run(":8080")
}
