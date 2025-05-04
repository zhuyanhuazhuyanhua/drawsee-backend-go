package handler

import (
	"drawsee/exception"
	"drawsee/pojo"
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// GetErrorResponse 根据 ApiError 创建错误响应
func GetErrorResponse(c *gin.Context, err exception.ApiError) {
	result := pojo.NewErrorResult(err.Code, err.Message)
	c.JSON(err.Code, result)
}

// HandleApiException 处理自定义异常
func HandleApiException(c *gin.Context, e *exception.ApiException) {
	GetErrorResponse(c, e.GetError())
}

// HandleValidationError 处理参数校验异常
func HandleValidationError(c *gin.Context) {
	apiError, _ := exception.GetApiError(exception.ParamError)
	GetErrorResponse(c, apiError)
}

// HandleNotLoginException 处理未登录异常
func HandleNotLoginException(c *gin.Context) {
	apiError, _ := exception.GetApiError(exception.NotLogin)
	GetErrorResponse(c, apiError)
}

// HandleException 处理其他异常
func HandleException(c *gin.Context, e error) {
	fmt.Println("Error:", e)
	apiError, _ := exception.GetApiError(exception.SystemError)
	GetErrorResponse(c, apiError)
}

// GlobalExceptionHandler 全局异常处理中间件
func GlobalExceptionHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// 处理自定义异常
		if err := c.Errors.Last(); err != nil {
			var validationErrs validator.ValidationErrors
			if errors.As(err.Err, &validationErrs) {
				HandleValidationError(c)
				return
			}
			switch e := err.Err.(type) {
			case *exception.ApiException:
				HandleApiException(c, e)
			case exception.NotLoginError:
				HandleNotLoginException(c)
			default:
				HandleException(c, e)
			}
		}
	}
}
