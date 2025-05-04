package config

import (
	"testing"
)

func TestPromptService(t *testing.T) {
	// 创建ResourceLoader实例
	resourceLoader := &FileResourceLoader{}

	// 创建代理实例
	promptService := NewPromptServiceProxy(resourceLoader)

	// 测试GetPrompt方法
	expected := "This is a prompt from the service."
	actual := promptService.GetPrompt()
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func TestFileResourceLoader(t *testing.T) {
	// 创建ResourceLoader实例
	resourceLoader := &FileResourceLoader{}

	// 测试GetResource方法
	expected := "Hello, World!"
	actual, err := resourceLoader.GetResource("C:/Users/1/Desktop/drawsee/resources/prompt/test.txt")
	if err != nil {
		t.Errorf("Error reading file: %v", err)
	}
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}
