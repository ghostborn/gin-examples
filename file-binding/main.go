package main

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type BindFile struct {
	Name  string                `form:"name" binding:"required"`
	Email string                `form:"email" binding:"required"`
	File  *multipart.FileHeader `form:"file" binding:"required"`
}

func main() {
	router := gin.Default()
	router.MaxMultipartMemory = 8 << 20

	router.Static("/", "./public")

	router.POST("/upload", func(ctx *gin.Context) {
		// 1. 初始化绑定结构体实例
		var bindFile BindFile
		// 2. 绑定并校验表单数据
		// ShouldBind 会自动根据请求类型（multipart/form-data）解析参数
		// 并执行 binding 标签的校验规则
		if err := ctx.ShouldBind(&bindFile); err != nil {
			// 绑定/校验失败：返回 400 错误
			ctx.String(http.StatusBadRequest, fmt.Sprintf("err:%s", err.Error()))
			return
		}
		// 3.提取上传文件并保存到本地
		file := bindFile.File // 获取文件元信息
		// 取文件名（剔除路径，仅保留文件名，避免路径注入攻击）
		dst := filepath.Base(file.Filename)
		// 保存文件到当前工作目录，文件名使用上传的原始文件名
		if err := ctx.SaveUploadedFile(file, dst); err != nil {
			// 保存失败：返回 400 错误
			ctx.String(http.StatusBadRequest, fmt.Sprintf("upload file err:%s", err.Error()))
			return
		}
		// 4. 上传成功：返回200成功信息
		ctx.String(
			http.StatusOK,
			fmt.Sprintf(
				"File %s uploaded successfully with fields name=%s and email=%s.",
				file.Filename,
				bindFile.Name,
				bindFile.Email,
			),
		)
	})
	_ = router.Run(":8080")
}
