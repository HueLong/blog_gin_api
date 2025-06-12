package router

import (
	"github.com/gin-gonic/gin"
	"github.com/huelong/blog_gin_api/internal/middleware"
)

// SetupRouter 配置路由
func SetupRouter() *gin.Engine {
	// 创建默认的 gin 路由引擎
	r := gin.Default()

	// 使用中间件
	r.Use(middleware.Logger())
	r.Use(middleware.Recovery())
	r.Use(middleware.Cors())

	// API 版本分组
	v1 := r.Group("/api/v1")
	{
		// 健康检查
		v1.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"status": "ok",
			})
		})

		// TODO: 在这里添加其他路由组
		// 例如：用户相关路由、文章相关路由等
	}

	return r
} 