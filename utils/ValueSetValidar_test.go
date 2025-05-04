package utils

import (
	"testing"
)

func TestValueSetValidator(t *testing.T) {
	validator := NewValueSetValidator([]string{"male", "female"})

	tests := []struct {
		name    string
		value   string
		wantErr bool
	}{
		{"Valid male", "male", false},
		{"Valid female", "female", false},
		{"Invalid value", "other", true},
		{"Empty value", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator.Validate(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil {
				t.Logf("Error: %v", err)
			}
		})
	}
}
