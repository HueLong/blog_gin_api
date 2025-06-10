package response

import (
	"blog_gin_api/internal/pkg/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response 标准响应结构
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// Success 成功响应
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    0,
		Message: "success",
		Data:    data,
	})
}

// Error 错误响应
func Error(c *gin.Context, err error) {
	if e, ok := err.(*errors.Error); ok {
		c.JSON(e.HTTPStatus(), Response{
			Code:    e.Code,
			Message: e.Message,
		})
		return
	}

	// 未知错误
	c.JSON(http.StatusInternalServerError, Response{
		Code:    errors.ErrInternalServer,
		Message: "Internal server error",
	})
}

// List 列表响应
type List struct {
	Total int64       `json:"total"`
	Items interface{} `json:"items"`
}

// ListResponse 返回列表数据
func ListResponse(c *gin.Context, total int64, items interface{}) {
	Success(c, List{
		Total: total,
		Items: items,
	})
}

// Page 分页参数
type Page struct {
	Page     int `form:"page" binding:"required,min=1"`
	PageSize int `form:"page_size" binding:"required,min=1,max=100"`
}

// GetOffset 获取偏移量
func (p *Page) GetOffset() int {
	return (p.Page - 1) * p.PageSize
}

// GetLimit 获取限制
func (p *Page) GetLimit() int {
	return p.PageSize
} 