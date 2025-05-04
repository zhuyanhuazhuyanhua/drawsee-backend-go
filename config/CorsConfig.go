package config

import (
	"net/http"

	"github.com/gorilla/mux"
)

// AddCorsMiddleware 为路由添加 CORS 中间件
func AddCorsMiddleware(router *mux.Router) *mux.Router {
	router.Use(corsMiddleware)
	return router
}

// corsMiddleware 实现 CORS 中间件逻辑
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 设置允许的源模式，这里使用通配符
		w.Header().Set("Access-Control-Allow-Origin", "*")
		// 设置允许的请求头，使用通配符
		w.Header().Set("Access-Control-Allow-Headers", "*")
		// 设置允许的请求方法，使用通配符
		w.Header().Set("Access-Control-Allow-Methods", "*")
		// 允许携带凭证
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		// 设置预检请求的缓存时间
		w.Header().Set("Access-Control-Max-Age", "168000")

		// 处理预检请求
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		// 继续处理其他请求
		next.ServeHTTP(w, r)
	})
}
