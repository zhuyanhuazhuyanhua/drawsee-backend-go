package utils

import (
	"errors"
)

// ValueSetValidator 用于验证值是否在允许的值集合中
type ValueSetValidator struct {
	allowedValues []string
}

// NewValueSetValidator 创建一个新的 ValueSetValidator 实例
func NewValueSetValidator(values []string) *ValueSetValidator {
	return &ValueSetValidator{
		allowedValues: values,
	}
}

// Validate 检查值是否在允许的值集合中
func (v *ValueSetValidator) Validate(value string) error {
	if value == "" {
		return errors.New("value cannot be empty")
	}
	for _, allowed := range v.allowedValues {
		if value == allowed {
			return nil
		}
	}
	return errors.New("value must be one of the specified values")
}
