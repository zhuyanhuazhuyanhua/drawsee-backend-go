package dto

import (
	"encoding/json"
	"testing"
)

func TestAdminRegisterDTO(t *testing.T) {
	// Test case 1: Valid AdminRegisterDTO
	userID := int64(123)
	dto := AdminRegisterDTO{UserID: &userID}

	// Marshal to JSON
	jsonData, err := json.Marshal(dto)
	if err != nil {
		t.Fatalf("Failed to marshal AdminRegisterDTO: %v", err)
	}

	// Unmarshal from JSON
	var newDTO AdminRegisterDTO
	err = json.Unmarshal(jsonData, &newDTO)
	if err != nil {
		t.Fatalf("Failed to unmarshal AdminRegisterDTO: %v", err)
	}

	// Validate
	if newDTO.UserID == nil || *newDTO.UserID != userID {
		t.Fatalf("Unmarshaled DTO does not match expected values")
	}

	// Test case 2: Invalid AdminRegisterDTO
	invalidDTO := AdminRegisterDTO{UserID: nil}
	if invalidDTO.UserID != nil {
		t.Fatalf("Expected UserID to be nil")
	}

	// Test case 3: JSON Marshaling and Unmarshaling
	expectedJSON := `{"userId":123}`
	actualJSON, _ := json.Marshal(dto)
	if string(actualJSON) != expectedJSON {
		t.Fatalf("Expected JSON: %s, got: %s", expectedJSON, actualJSON)
	}

	var unmarshaledDTO AdminRegisterDTO
	err = json.Unmarshal([]byte(expectedJSON), &unmarshaledDTO)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
	}
	if unmarshaledDTO.UserID == nil || *unmarshaledDTO.UserID != userID {
		t.Fatalf("Unmarshaled DTO does not match expected values")
	}
}
