package annotation

import (
	"reflect"
	"testing"
)

// 定义一个带有 RespondToQuery 方法的结构体
type testStruct struct{}

// 定义 RespondToQuery 方法
func (ts *testStruct) RespondToQuery(prompt PromptResource, query string) {
	// 方法实现可以为空，只是为了测试
}

// TestGetPromptResource 测试 GetPromptResource 函数
func TestGetPromptResource(t *testing.T) {
	// 创建一个测试结构体
	test := &testStruct{}

	// 获取方法的反射信息
	method, found := reflect.TypeOf(test).MethodByName("RespondToQuery")
	if !found {
		t.Fatalf("Expected method to be found, but it was not")
	}

	// 调用 GetPromptResource 函数
	promptResource, err := GetPromptResource(method)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if promptResource == nil {
		t.Fatalf("Expected PromptResource, got nil")
	}

	expected := &PromptResource{}
	if promptResource.FromResource != expected.FromResource {
		t.Errorf("Expected FromResource to be %s, got %s", expected.FromResource, promptResource.FromResource)
	}
}
