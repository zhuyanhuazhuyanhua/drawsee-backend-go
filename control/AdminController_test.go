package control

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAdminController(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// 创建一个测试的 Gin 引擎
	engine := gin.New()

	// 注册路由
	RegisterRoutes(engine)

	// 测试 Register
	t.Run("Test Register", func(t *testing.T) {
		w := httptest.NewRecorder()
		req := httptest.NewRequest("POST", "/admin/register", strings.NewReader(`{"username":"test","password":"password"}`))
		req.Header.Set("Content-Type", "application/json")
		engine.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
	})

	// 测试 Login
	t.Run("Test Login", func(t *testing.T) {
		w := httptest.NewRecorder()
		req := httptest.NewRequest("POST", "/admin/login", strings.NewReader(`{"username":"test","password":"password"}`))
		req.Header.Set("Content-Type", "application/json")
		engine.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
	})

	// 测试 CheckLogin
	t.Run("Test CheckLogin", func(t *testing.T) {
		w := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/admin/check_login", nil)
		engine.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
	})

	// 测试 GetInvitationCodesByPage
	t.Run("Test GetInvitationCodesByPage", func(t *testing.T) {
		w := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/admin/invitation_codes?page=1&size=10", nil)
		engine.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
	})

	// 测试 CreateInvitationCode
	t.Run("Test CreateInvitationCode", func(t *testing.T) {
		w := httptest.NewRecorder()
		req := httptest.NewRequest("POST", "/admin/invitation_codes", strings.NewReader(`{"code":"12345"}`))
		req.Header.Set("Content-Type", "application/json")
		engine.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
	})

	// 测试 SendInvitationCode
	t.Run("Test SendInvitationCode", func(t *testing.T) {
		w := httptest.NewRecorder()
		req := httptest.NewRequest("POST", "/admin/invitation_codes/1", strings.NewReader(`{"email":"test@example.com"}`))
		req.Header.Set("Content-Type", "application/json")
		engine.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
	})

	// 测试 GetStatistics
	t.Run("Test GetStatistics", func(t *testing.T) {
		w := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/admin/statistics", nil)
		engine.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
	})
}
