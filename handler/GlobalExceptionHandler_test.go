package handler_test

import (
	"bytes"
	"drawsee/exception"
	"drawsee/handler"
	"drawsee/pojo"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

// 初始化验证器
var validate = validator.New()

func TestGlobalExceptionHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	// 初始化验证器
	gin.ForceConsoleColor()
	router := gin.Default()
	router.Use(handler.GlobalExceptionHandler())

	// 测试自定义异常
	router.GET("/api_exception", func(c *gin.Context) {
		apiError, _ := exception.GetApiError(exception.SystemError)
		apiException := exception.NewApiException(apiError)
		c.Error(apiException)
	})

	// 测试参数校验异常
	router.POST("/validation_error", func(c *gin.Context) {
		var data struct {
			Name string `json:"name" validate:"required"`
		}
		if err := c.ShouldBindJSON(&data); err != nil {
			c.Error(err)
			c.Abort()
			return
		}
		// 显式调用验证逻辑
		if err := validate.Struct(data); err != nil {
			c.Error(err)
			c.Abort()
			return
		}
	})

	// 测试未登录异常
	router.GET("/not_login", func(c *gin.Context) {
		c.Error(exception.NotLoginError{})
	})

	// 测试其他异常
	router.GET("/other_exception", func(c *gin.Context) {
		c.Error(fmt.Errorf("some other error"))
	})

	// 测试自定义异常
	t.Run("Test ApiException", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/api_exception", nil)
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusInternalServerError, resp.Code)
		var result pojo.Result[any]
		json.Unmarshal(resp.Body.Bytes(), &result)
		assert.Equal(t, 500, result.Code)
		assert.Equal(t, "服务器内部错误", result.Message)
	})

	// 测试参数校验异常
	t.Run("Test ValidationError", func(t *testing.T) {
		reqBody := []byte(`{"name": ""}`)
		req := httptest.NewRequest("POST", "/validation_error", bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusBadRequest, resp.Code)
		var result pojo.Result[any]
		json.Unmarshal(resp.Body.Bytes(), &result)
		assert.Equal(t, 400, result.Code)
		assert.Equal(t, "参数错误", result.Message)
	})

	// 测试未登录异常
	t.Run("Test NotLoginException", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/not_login", nil)
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusUnauthorized, resp.Code)
		var result pojo.Result[any]
		json.Unmarshal(resp.Body.Bytes(), &result)
		assert.Equal(t, 401, result.Code)
		assert.Equal(t, "未登录", result.Message)
	})

	// 测试其他异常
	t.Run("Test OtherException", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/other_exception", nil)
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusInternalServerError, resp.Code)
		var result pojo.Result[any]
		json.Unmarshal(resp.Body.Bytes(), &result)
		assert.Equal(t, 500, result.Code)
		assert.Equal(t, "服务器内部错误", result.Message)
	})
}
