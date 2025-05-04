package entity

import (
	"encoding/json"
	"testing"
)

func TestAdmin(t *testing.T) {
	// Test NewAdmin
	admin := NewAdmin(123)
	if admin.UserID != 123 {
		t.Errorf("NewAdmin failed, expected UserID to be 123, got %v", admin.UserID)
	}

	// Test String method
	expectedString := "Admin{ID: <nil>, UserID: 123}"
	if admin.String() != expectedString {
		t.Errorf("String method failed, expected %v, got %v", expectedString, admin.String())
	}

	// Test MarshalJSON
	jsonData, err := json.Marshal(admin)
	if err != nil {
		t.Errorf("MarshalJSON failed: %v", err)
	}
	expectedJSON := `{"id":null,"userId":123}`
	if string(jsonData) != expectedJSON {
		t.Errorf("MarshalJSON failed, expected %v, got %v", expectedJSON, string(jsonData))
	}

	// Test UnmarshalJSON
	var newAdmin Admin
	err = json.Unmarshal(jsonData, &newAdmin)
	if err != nil {
		t.Errorf("UnmarshalJSON failed: %v", err)
	}
	if newAdmin.UserID != 123 {
		t.Errorf("UnmarshalJSON failed, expected UserID to be 123, got %v", newAdmin.UserID)
	}
}
