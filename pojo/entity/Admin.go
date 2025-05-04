package entity

import (
	"encoding/json"
	"fmt"
)

// Admin represents an admin entity
type Admin struct {
	ID     *int64 `json:"id"`
	UserID int64  `json:"userId"`
}

// NewAdmin creates a new Admin instance with the given userId
func NewAdmin(userID int64) *Admin {
	return &Admin{
		UserID: userID,
	}
}

// String returns a string representation of the Admin instance
func (a *Admin) String() string {
	return fmt.Sprintf("Admin{ID: %v, UserID: %v}", a.ID, a.UserID)
}

// MarshalJSON implements json.Marshaler interface for Admin
func (a *Admin) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ID     *int64 `json:"id"`
		UserID int64  `json:"userId"`
	}{
		ID:     a.ID,
		UserID: a.UserID,
	})
}

// UnmarshalJSON implements json.Unmarshaler interface for Admin
func (a *Admin) UnmarshalJSON(data []byte) error {
	var admin struct {
		ID     *int64 `json:"id"`
		UserID int64  `json:"userId"`
	}
	if err := json.Unmarshal(data, &admin); err != nil {
		return err
	}
	a.ID = admin.ID
	a.UserID = admin.UserID
	return nil
}
