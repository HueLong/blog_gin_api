package main

import (
	"blog_gin_api/internal/middleware"
	"blog_gin_api/internal/pkg/config"
	"blog_gin_api/internal/pkg/logger"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	// 初始化配置
	if err := config.Init(); err != nil {
		panic(fmt.Sprintf("init config failed: %v", err))
	}

	// 初始化日志
	if err := logger.Init(); err != nil {
		panic(fmt.Sprintf("init logger failed: %v", err))
	}

	// 设置运行模式
	gin.SetMode(config.GlobalConfig.Server.Mode)

	// 创建 Gin 引擎
	r := gin.New()

	// 使用中间件
	r.Use(middleware.Logger())
	r.Use(middleware.Recovery())
	r.Use(middleware.CORS())

	// 注册路由
	registerRoutes(r)

	// 创建 HTTP 服务器
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", config.GlobalConfig.Server.Port),
		Handler:      r,
		ReadTimeout:  time.Duration(config.GlobalConfig.Server.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(config.GlobalConfig.Server.WriteTimeout) * time.Second,
	}

	// 在独立的 goroutine 中启动服务器
	go func() {
		logger.Info("server is starting", zap.String("addr", srv.Addr))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("listen failed", zap.Error(err))
		}
	}()

	// 等待中断信号以优雅地关闭服务器
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info("shutting down server...")

	// 设置 5 秒的超时时间
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatal("server forced to shutdown", zap.Error(err))
	}

	logger.Info("server exiting")
}

func registerRoutes(r *gin.Engine) {
	// 健康检查
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// API 版本分组
	v1 := r.Group("/api/v1")
	{
		// TODO: 添加更多路由
		v1.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Welcome to Blog API v1",
			})
		})
	}
} 