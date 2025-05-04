package handler_test

import (
	"drawsee/handler"
	"drawsee/pojo"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGlobalResponseHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	// 使用全局响应处理中间件
	router.Use(handler.GlobalResponseHandler())

	// 测试非 Result 类型响应
	router.GET("/non-result", func(c *gin.Context) {
		c.JSON(200, "test data")
	})

	// 测试 Result 类型响应
	router.GET("/result", func(c *gin.Context) {
		result := pojo.NewResult("test data")
		c.JSON(200, result)
	})

	// 测试非 Result 类型响应
	t.Run("Test Non-Result Response", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/non-result", nil)
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusOK, resp.Code)
		var result pojo.Result[interface{}]
		err := json.Unmarshal(resp.Body.Bytes(), &result)
		assert.NoError(t, err)
		assert.Equal(t, 200, result.Code)
		assert.Equal(t, "success", result.Message)
		assert.Equal(t, "test data", result.Data)
	})

	// 测试 Result 类型响应
	t.Run("Test Result Response", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/result", nil)
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusOK, resp.Code)
		var result pojo.Result[interface{}]
		err := json.Unmarshal(resp.Body.Bytes(), &result)
		assert.NoError(t, err)
		assert.Equal(t, 200, result.Code)
		assert.Equal(t, "success", result.Message)
		// 正确提取 Data 字段并比较
		assert.Equal(t, "test data", result.Data)
	})
}
