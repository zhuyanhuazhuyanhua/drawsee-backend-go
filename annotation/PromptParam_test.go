package annotation

import (
	"fmt"
	"testing"
)

// TestStruct 用于测试的结构体
type TestStruct struct {
	Param string `promptparam:"test_value"`
}

func TestGetPromptParamValue(t *testing.T) {
	testStruct := &TestStruct{}
	value, err := GetPromptParamValue(testStruct)
	if err != nil {
		t.Fatalf("Failed to get PromptParam value: %v", err)
	}
	if value != "test_value" {
		t.Errorf("Expected 'test_value', got '%s'", value)
	} else {
		fmt.Println("Test passed:", value)
	}
}

func TestInvalidInput(t *testing.T) {
	// 测试非结构体指针
	value, err := GetPromptParamValue("invalid input")
	if err == nil {
		t.Errorf("Expected error for invalid input, got value: %s", value)
	} else {
		fmt.Println("Test passed:", err)
	}

	// 测试没有PromptParam标签的结构体
	type NoTagStruct struct {
		Param string
	}
	noTagStruct := &NoTagStruct{}
	value, err = GetPromptParamValue(noTagStruct)
	if err == nil {
		t.Errorf("Expected error for struct without PromptParam tag, got value: %s", value)
	} else {
		fmt.Println("Test passed:", err)
	}
}
