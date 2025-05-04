package pojo

import (
	"encoding/json"
	"fmt"
	"time"
)

// Result 泛型结果结构体
type Result[T any] struct {
	Code      int       `json:"code"`
	Message   string    `json:"message"`
	Data      T         `json:"data"`
	Timestamp time.Time `json:"timestamp"`
}

// NewResult 创建一个新的 Result 实例
func NewResult[T any](data T) *Result[T] {
	return &Result[T]{
		Code:      200,
		Message:   "success",
		Data:      data,
		Timestamp: time.Now(),
	}
}

// NewErrorResult 创建一个新的错误 Result 实例
func NewErrorResult(code int, message string) *Result[any] {
	return &Result[any]{
		Code:      code,
		Message:   message,
		Data:      nil,
		Timestamp: time.Now(),
	}
}

// String 返回 Result 的字符串表示
func (r *Result[T]) String() string {
	data, err := json.Marshal(r)
	if err != nil {
		return fmt.Sprintf("Error marshaling result: %v", err)
	}
	return string(data)
}
