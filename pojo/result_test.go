package pojo

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestResult(t *testing.T) {
	// 测试成功结果
	successData := map[string]string{"key": "value"}
	successResult := NewResult(successData)
	if successResult.Code != 200 {
		t.Errorf("Expected code 200, got %d", successResult.Code)
	}
	if successResult.Message != "success" {
		t.Errorf("Expected message 'success', got '%s'", successResult.Message)
	}
	// 由于不能直接比较 map，我们通过 JSON 序列化来比较
	successDataBytes, _ := json.Marshal(successData)
	successResultDataBytes, _ := json.Marshal(successResult.Data)
	if string(successDataBytes) != string(successResultDataBytes) {
		t.Errorf("Expected data %v, got %v", successData, successResult.Data)
	}
	if successResult.Timestamp.IsZero() {
		t.Error("Expected non-zero timestamp")
	}

	// 测试错误结果
	errorResult := NewErrorResult(404, "Not Found")
	if errorResult.Code != 404 {
		t.Errorf("Expected code 404, got %d", errorResult.Code)
	}
	if errorResult.Message != "Not Found" {
		t.Errorf("Expected message 'Not Found', got '%s'", errorResult.Message)
	}
	if errorResult.Data != nil {
		t.Errorf("Expected nil data, got %v", errorResult.Data)
	}
	if errorResult.Timestamp.IsZero() {
		t.Error("Expected non-zero timestamp")
	}

	// 测试 String 方法
	successResultString := successResult.String()
	successResultJSON := map[string]interface{}{}
	if err := json.Unmarshal([]byte(successResultString), &successResultJSON); err != nil {
		t.Errorf("Failed to unmarshal success result JSON: %v", err)
	}
	if successResultJSON["code"] != float64(200) {
		t.Errorf("Expected code 200, got %v", successResultJSON["code"])
	}
	if successResultJSON["message"] != "success" {
		t.Errorf("Expected message 'success', got '%s'", successResultJSON["message"])
	}
	if successResultJSON["data"].(map[string]interface{})["key"] != "value" {
		t.Errorf("Expected data key 'value', got %v", successResultJSON["data"].(map[string]interface{})["key"])
	}

	errorResultString := errorResult.String()
	errorResultJSON := map[string]interface{}{}
	if err := json.Unmarshal([]byte(errorResultString), &errorResultJSON); err != nil {
		t.Errorf("Failed to unmarshal error result JSON: %v", err)
	}
	if errorResultJSON["code"] != float64(404) {
		t.Errorf("Expected code 404, got %v", errorResultJSON["code"])
	}
	if errorResultJSON["message"] != "Not Found" {
		t.Errorf("Expected message 'Not Found', got '%s'", errorResultJSON["message"])
	}
	if errorResultJSON["data"] != nil {
		t.Errorf("Expected nil data, got %v", errorResultJSON["data"])
	}
}

func ExampleResult() {
	// 示例：创建成功结果
	successData := map[string]string{"key": "value"}
	successResult := NewResult(successData)
	fmt.Println(successResult.String())

	// 示例：创建错误结果
	errorResult := NewErrorResult(404, "Not Found")
	fmt.Println(errorResult.String())

	// 输出：
	// {"code":200,"message":"success","data":{"key":"value"},"timestamp":"2025-05-04T15:19:00Z"}
	// {"code":404,"message":"Not Found","data":null,"timestamp":"2025-05-04T15:19:00Z"}
}
