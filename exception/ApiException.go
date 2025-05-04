package exception

import (
	"errors"
	"fmt"
)

// ApiException 定义了API异常类型
type ApiException struct {
	error ApiError
}

// NewApiException 创建一个新的ApiException
func NewApiException(err ApiError) *ApiException {
	return &ApiException{error: err}
}

// Error 实现error接口
func (e *ApiException) Error() string {
	return fmt.Sprintf("code: %d, message: %s", e.error.Code, e.error.Message)
}

// Unwrap 实现errors.Unwrap接口，用于支持多层错误包装
func (e *ApiException) Unwrap() error {
	return errors.New(e.error.Message)
}

// GetError 获取 ApiException 内部的 ApiError 实例
func (e *ApiException) GetError() ApiError {
	return e.error
}
