package annotation

import (
	"fmt"
	"reflect"
)

// PromptParam 是一个模拟Java注解的结构体
type PromptParam struct {
	Value string `json:"value"`
}

// GetPromptParamValue 通过反射获取参数上的PromptParam标签值
func GetPromptParamValue(param interface{}) (string, error) {
	val := reflect.ValueOf(param)
	if val.Kind() != reflect.Ptr || val.Elem().Kind() != reflect.Struct {
		return "", fmt.Errorf("param must be a pointer to a struct")
	}

	elem := val.Elem()
	for i := 0; i < elem.NumField(); i++ {
		field := elem.Type().Field(i)
		tag := field.Tag.Get("promptparam")
		if tag != "" {
			return tag, nil
		}
	}

	return "", fmt.Errorf("no PromptParam tag found")
}
