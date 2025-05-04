package config

import (
	"log"
	"net/http"
)

// SaTokenMiddleware 模拟 SaToken 的中间件
func SaTokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 在这里添加 SaToken 的逻辑
		log.Println("SaToken Middleware: Checking token...")
		// 假设这里检查了 token，如果没有问题，继续处理请求
		next.ServeHTTP(w, r)
	})
}

// SetupRoutes 设置路由和中间件
func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	// 添加中间件
	mux.Handle("/", SaTokenMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})))

	return mux
}
