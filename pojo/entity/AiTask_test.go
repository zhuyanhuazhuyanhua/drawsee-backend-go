package entity

import (
	"encoding/json"
	"testing"
	"time"
)

func TestAiTask(t *testing.T) {
	// Test NewAiTask
	aiTask := NewAiTask("type1", "data1", "status1", 123, 456)
	if aiTask.Type != "type1" || aiTask.Data != "data1" || aiTask.Status != "status1" || aiTask.UserID != 123 || aiTask.ConvID != 456 {
		t.Errorf("NewAiTask failed, expected type1, data1, status1, 123, 456, got %v, %v, %v, %v, %v",
			aiTask.Type, aiTask.Data, aiTask.Status, aiTask.UserID, aiTask.ConvID)
	}

	// Test String method
	expectedString := "AiTask{ID: <nil>, Type: type1, Data: data1, Result: , Status: status1, Tokens: <nil>, UserID: 123, ConvID: 456, CreatedAt: <nil>, UpdatedAt: <nil>, IsDeleted: <nil>}"
	if aiTask.String() != expectedString {
		t.Errorf("String method failed, expected %v, got %v", expectedString, aiTask.String())
	}

	// Test MarshalJSON
	jsonData, err := json.Marshal(aiTask)
	if err != nil {
		t.Errorf("MarshalJSON failed: %v", err)
	}
	expectedJSON := `{"id":null,"type":"type1","data":"data1","result":"","status":"status1","tokens":null,"userId":123,"convId":456,"createdAt":null,"updatedAt":null,"isDeleted":null}`
	if string(jsonData) != expectedJSON {
		t.Errorf("MarshalJSON failed, expected %v, got %v", expectedJSON, string(jsonData))
	}

	// Test UnmarshalJSON
	var newAiTask AiTask
	err = json.Unmarshal(jsonData, &newAiTask)
	if err != nil {
		t.Errorf("UnmarshalJSON failed: %v", err)
	}
	if newAiTask.Type != "type1" || newAiTask.Data != "data1" || newAiTask.Status != "status1" || newAiTask.UserID != 123 || newAiTask.ConvID != 456 {
		t.Errorf("UnmarshalJSON failed, expected type1, data1, status1, 123, 456, got %v, %v, %v, %v, %v",
			newAiTask.Type, newAiTask.Data, newAiTask.Status, newAiTask.UserID, newAiTask.ConvID)
	}

	// Test with additional fields
	createdAt := time.Date(2025, 1, 29, 17, 8, 0, 0, time.UTC)
	updatedAt := time.Date(2025, 1, 29, 17, 9, 0, 0, time.UTC)
	isDeleted := true

	// Create pointers for int64 and bool values
	var idValue int64 = 1
	var tokensValue int64 = 100
	var isDeletedValue bool = true

	aiTaskWithFields := &AiTask{
		ID:        &idValue,
		Type:      "type2",
		Data:      "data2",
		Result:    "result2",
		Status:    "status2",
		Tokens:    &tokensValue,
		UserID:    789,
		ConvID:    1011,
		CreatedAt: &createdAt,
		UpdatedAt: &updatedAt,
		IsDeleted: &isDeletedValue,
	}
	jsonDataWithFields, err := json.Marshal(aiTaskWithFields)
	if err != nil {
		t.Errorf("MarshalJSON failed: %v", err)
	}
	var newAiTaskWithFields AiTask
	err = json.Unmarshal(jsonDataWithFields, &newAiTaskWithFields)
	if err != nil {
		t.Errorf("UnmarshalJSON failed: %v", err)
	}
	if newAiTaskWithFields.Type != "type2" || newAiTaskWithFields.Data != "data2" || newAiTaskWithFields.Status != "status2" || newAiTaskWithFields.UserID != 789 || newAiTaskWithFields.ConvID != 1011 {
		t.Errorf("UnmarshalJSON failed, expected type2, data2, status2, 789, 1011, got %v, %v, %v, %v, %v",
			newAiTaskWithFields.Type, newAiTaskWithFields.Data, newAiTaskWithFields.Status, newAiTaskWithFields.UserID, newAiTaskWithFields.ConvID)
	}
	if newAiTaskWithFields.Tokens == nil || *newAiTaskWithFields.Tokens != 100 {
		t.Errorf("UnmarshalJSON failed, expected tokens 100, got %v", newAiTaskWithFields.Tokens)
	}
	if newAiTaskWithFields.CreatedAt == nil || newAiTaskWithFields.CreatedAt.String() != createdAt.String() {
		t.Errorf("UnmarshalJSON failed, expected createdAt %v, got %v", createdAt.String(), newAiTaskWithFields.CreatedAt.String())
	}
	if newAiTaskWithFields.UpdatedAt == nil || newAiTaskWithFields.UpdatedAt.String() != updatedAt.String() {
		t.Errorf("UnmarshalJSON failed, expected updatedAt %v, got %v", updatedAt.String(), newAiTaskWithFields.UpdatedAt.String())
	}
	if newAiTaskWithFields.IsDeleted == nil || *newAiTaskWithFields.IsDeleted != isDeleted {
		t.Errorf("UnmarshalJSON failed, expected isDeleted %v, got %v", isDeleted, newAiTaskWithFields.IsDeleted)
	}
}
