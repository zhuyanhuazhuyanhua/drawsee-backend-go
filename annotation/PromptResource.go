package annotation

import (
	"fmt"
	"reflect"
)

// PromptResource 用于标记方法的资源来源
type PromptResource struct {
	FromResource string
}

// GetPromptResource 通过反射获取方法上的PromptResource
func GetPromptResource(method reflect.Method) (*PromptResource, error) {
	// 检查方法是否有一个PromptResource类型的参数
	for i := 0; i < method.Type.NumIn(); i++ {
		param := method.Type.In(i)
		if param.Kind() == reflect.Struct && param.Name() == "PromptResource" {
			return &PromptResource{}, nil
		}
	}

	return nil, fmt.Errorf("no PromptResource found in method parameters")
}
