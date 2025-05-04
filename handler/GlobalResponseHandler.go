package handler

import (
	"bytes"
	"drawsee/pojo"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

// responseBodyWriter is a custom response writer to capture the response body
type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r *responseBodyWriter) Write(b []byte) (int, error) {
	return r.body.Write(b)
}

// globalResponseHandler 全局响应处理中间件
func globalResponseHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Create a custom response writer
		rbw := &responseBodyWriter{
			ResponseWriter: c.Writer,
			body:           bytes.NewBuffer(nil),
		}
		c.Writer = rbw

		// 执行后续处理逻辑
		c.Next()

		// 获取响应状态码
		statusCode := rbw.Status()
		if statusCode >= 200 && statusCode < 300 {
			// 获取响应体
			bodyBytes := rbw.body.Bytes()
			var body interface{}
			if len(bodyBytes) > 0 {
				// 尝试将响应体解析为 Result 类型
				var result pojo.Result[interface{}]
				if err := json.Unmarshal(bodyBytes, &result); err == nil {
					// 如果解析成功，说明本身就是 Result 类型，直接返回
					rbw.ResponseWriter.Header().Set("Content-Type", "application/json")
					rbw.ResponseWriter.WriteHeader(statusCode)
					_, err := rbw.ResponseWriter.Write(bodyBytes)
					if err != nil {
						c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
							"error": "Failed to write response",
						})
					}
					return
				}
				// 解析失败，按普通响应处理
				if err := json.Unmarshal(bodyBytes, &body); err != nil {
					body = string(bodyBytes)
				}

				// Wrap it into a Result type
				newResult := pojo.NewResult(body)
				newResultBytes, err := json.Marshal(newResult)
				if err != nil {
					c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
						"error": "Failed to marshal response",
					})
					return
				}

				// Reset the response writer
				rbw.ResponseWriter.Header().Set("Content-Type", "application/json")
				rbw.ResponseWriter.WriteHeader(http.StatusOK)
				rbw.body.Reset()

				// Write the new response
				_, err = rbw.ResponseWriter.Write(newResultBytes)
				if err != nil {
					c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
						"error": "Failed to write response",
					})
				}
			}
		}
	}
}

// GlobalResponseHandler 对外暴露的全局响应处理中间件函数
func GlobalResponseHandler() gin.HandlerFunc {
	return globalResponseHandler()
}
