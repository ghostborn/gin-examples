package main

import (
	"context"
	"crypto/rand"
	"embed"
	"encoding/base64"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

var (
	clientId     = os.Getenv("GITHUB_CLIENT_ID")
	clientSecret = os.Getenv("GITHUB_CLIENT_SECRET")
	redirectURL  = "http://localhost:8080/callback"
)

var (
	stateCache = cache.New(10*time.Minute, 20*time.Minute)

	oauth2Config = oauth2.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Endpoint:     github.Endpoint,
		Scopes:       []string{"user:email", "read:user"},
	}
	// GitHub 用户信息 API 地址
	userInfoURL = "https://api.github.com/user"
)

func generateRandomState() (string, error) {
	b := make([]byte, 32) // 生成32字节随机数
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	// 转成 URL 安全的 Base64 字符串（避免特殊字符）
	return base64.URLEncoding.EncodeToString(b), nil
}

type GitHubUser struct {
	ID        int    `json:"id"`
	Login     string `json:"login"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	AvatarURL string `json:"avatar_url"`
}

//go:embed templates/*
var templatesFS embed.FS

func main() {
	r := gin.Default()
	// 加载嵌入式模板
	r.SetHTMLTemplate(template.Must(template.ParseFS(templatesFS, "templates/*")))

	// 1. 首页路由：渲染登录入口页面
	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Github OAuth Example",
		})
	})

	// 2. 登录路由：重定向到 GitHub 授权页面
	r.GET("/login", func(ctx *gin.Context) {
		// 生成随机 state
		state, err := generateRandomState()
		if err != nil {
			ctx.String(http.StatusInternalServerError, "Unable to generate state value")
			return
		}

		// 存储 state 到缓存（用于回调时验证）
		stateCache.Set(state, true, cache.DefaultExpiration)

		// 生成 GitHub 授权 URL，并重定向用户到该地址
		authURL := oauth2Config.AuthCodeURL(state)
		ctx.Redirect(http.StatusFound, authURL)
	})

	r.GET("/callback", func(ctx *gin.Context) {
		// 步骤1：验证 state（防 CSRF）
		state := ctx.Query("state")
		if _, exists := stateCache.Get(state); !exists {
			ctx.String(http.StatusBadRequest, "Invalid state value")
			return
		}
		stateCache.Delete(state) // 验证后删除，避免重复使用

		// 步骤2：获取授权码 code
		code := ctx.Query("code")
		if code == "" {
			ctx.String(http.StatusBadRequest, "Authorization code not provided")
			return
		}

		// 步骤3：用 code 换取 access token
		token, err := oauth2Config.Exchange(context.Background(), code)
		if err != nil {
			ctx.String(
				http.StatusInternalServerError,
				"Unable to exchange access token: "+err.Error(),
			)
			return
		}

		// 步骤4：用 access token 调用 GitHub API 获取用户信息
		client := oauth2Config.Client(context.Background(), token) // 自动携带 token 的 HTTP 客户端
		resp, err := client.Get(userInfoURL)
		if err != nil {
			ctx.String(
				http.StatusInternalServerError,
				"Unable to retrieve user information: "+err.Error(),
			)
			return
		}

		defer resp.Body.Close() //确保响应体关闭

		// 解析用户信息 JSON
		var user GitHubUser
		if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
			ctx.String(
				http.StatusInternalServerError,
				"Unable to parse user information: "+err.Error(),
			)
			return
		}

		// 步骤5：渲染成功页面，展示用户信息
		ctx.HTML(http.StatusOK, "success.html", gin.H{
			"title":      "Authentication Successful",
			"username":   user.Login,
			"name":       user.Name,
			"email":      user.Email,
			"avatar_url": user.AvatarURL,
		})
	})

	// 4. 受保护路由示例（实际需校验登录态，此处简化）
	r.GET("/protected", func(ctx *gin.Context) {
		// 真实场景中：需校验 session/JWT token，确认用户已登录
		ctx.String(http.StatusOK, "This is a protected resource!")
	})

	// 启动服务
	log.Println("Server running at http://localhost:8080")
	r.Run(":8080")
}
