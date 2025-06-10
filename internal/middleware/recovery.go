package middleware

import (
	"blog_gin_api/internal/pkg/logger"
	"blog_gin_api/internal/pkg/response"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Recovery 恢复中间件
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// 记录堆栈信息
				logger.Error("panic recovered",
					zap.Any("error", err),
					zap.String("stack", string(debug.Stack())),
				)

				// 返回 500 错误
				c.AbortWithStatusJSON(http.StatusInternalServerError, response.Response{
					Code:    500,
					Message: "Internal server error",
				})
			}
		}()
		c.Next()
	}
} 