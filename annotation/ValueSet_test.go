package annotation

import (
	"testing"
)

func TestValueSet(t *testing.T) {
	// 创建一个 ValueSet 实例
	valueSet := NewValueSet(
		"Value must be one of the specified values",
		[]string{"group1", "group2"},
		[]string{"payload1", "payload2"},
		[]string{"male", "female"},
	)

	// 检查字段值
	if valueSet.Message != "Value must be one of the specified values" {
		t.Errorf("Expected message to be 'Value must be one of the specified values', got '%s'", valueSet.Message)
	}
	if len(valueSet.Groups) != 2 || valueSet.Groups[0] != "group1" || valueSet.Groups[1] != "group2" {
		t.Errorf("Expected groups to be ['group1', 'group2'], got %v", valueSet.Groups)
	}
	if len(valueSet.Payload) != 2 || valueSet.Payload[0] != "payload1" || valueSet.Payload[1] != "payload2" {
		t.Errorf("Expected payload to be ['payload1', 'payload2'], got %v", valueSet.Payload)
	}
	if len(valueSet.Values) != 2 || valueSet.Values[0] != "male" || valueSet.Values[1] != "female" {
		t.Errorf("Expected values to be ['male', 'female'], got %v", valueSet.Values)
	}
}
