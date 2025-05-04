package config

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSaTokenMiddleware(t *testing.T) {
	// 设置路由和中间件
	mux := SetupRoutes()

	// 创建一个测试请求
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("Error creating request: %v", err)
	}

	// 创建一个响应记录器
	rr := httptest.NewRecorder()

	// 调用中间件处理请求
	mux.ServeHTTP(rr, req)

	// 检查响应状态
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, rr.Code)
	}

	// 检查响应内容
	expected := "Hello, World!"
	if rr.Body.String() != expected {
		t.Errorf("Expected body %s, got %s", expected, rr.Body.String())
	}
}
