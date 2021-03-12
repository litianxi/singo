package server

import (
	"os"
	"singo/api"
	"singo/middleware"

	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 中间件, 顺序不能改
	r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	r.Use(middleware.Cors())
	r.Use(middleware.CurrentUser())

	// 路由
	v1 := r.Group("/api/v1")
	{
		v1.POST("ping", api.Ping)

		// 用户登录
		v1.POST("user/register", api.UserRegister)

		// 用户登录
		v1.POST("user/login", api.UserLogin)

		// 需要登录保护的
		auth := v1.Group("")
		auth.Use(middleware.AuthRequired())
		{
			// User Routing
			auth.GET("user/me", api.UserMe)
			auth.DELETE("user/logout", api.UserLogout)
		}
	}
	{
		// 短域名创建
		v1.POST("shortname", api.CreateShortName)

		// 短域名详情
		v1.GET("shortname/:id", api.ShortNameShow)

		// 短域名列表
		v1.GET("shortname", api.ListShortName)

		// 短域名删除
		v1.DELETE("shortname/:id", api.DeleteShortName)

		// 短域名修改
		v1.PUT("shortname/:id", api.UpdateShortName)
	}

	// swagger文档，浏览器打开 http://localhost:3000/swagger/index.html
	{
		r.StaticFile("/swagger.json", "./swagger/swagger.json")
		r.Static("/swagger", "./swagger/dist")
	}
	return r
}
