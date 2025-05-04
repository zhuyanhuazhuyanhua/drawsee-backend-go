package entity

import (
	"encoding/json"
	"fmt"
	"time"
)

// AiTask represents an AI task entity
type AiTask struct {
	ID        *int64     `json:"id"`
	Type      string     `json:"type"`
	Data      string     `json:"data"`
	Result    string     `json:"result"`
	Status    string     `json:"status"`
	Tokens    *int64     `json:"tokens"`
	UserID    int64      `json:"userId"`
	ConvID    int64      `json:"convId"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
	IsDeleted *bool      `json:"isDeleted"`
}

// NewAiTask creates a new AiTask instance with the given parameters
func NewAiTask(typeStr, data, status string, userID, convID int64) *AiTask {
	return &AiTask{
		Type:   typeStr,
		Data:   data,
		Status: status,
		UserID: userID,
		ConvID: convID,
	}
}

// String returns a string representation of the AiTask instance
func (a *AiTask) String() string {
	return fmt.Sprintf(
		"AiTask{ID: %v, Type: %v, Data: %v, Result: %v, Status: %v, Tokens: %v, UserID: %v, ConvID: %v, CreatedAt: %v, UpdatedAt: %v, IsDeleted: %v}",
		a.ID, a.Type, a.Data, a.Result, a.Status, a.Tokens, a.UserID, a.ConvID, a.CreatedAt, a.UpdatedAt, a.IsDeleted,
	)
}

// MarshalJSON implements json.Marshaler interface for AiTask
func (a *AiTask) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ID        *int64     `json:"id"`
		Type      string     `json:"type"`
		Data      string     `json:"data"`
		Result    string     `json:"result"`
		Status    string     `json:"status"`
		Tokens    *int64     `json:"tokens"`
		UserID    int64      `json:"userId"`
		ConvID    int64      `json:"convId"`
		CreatedAt *time.Time `json:"createdAt"`
		UpdatedAt *time.Time `json:"updatedAt"`
		IsDeleted *bool      `json:"isDeleted"`
	}{
		ID:        a.ID,
		Type:      a.Type,
		Data:      a.Data,
		Result:    a.Result,
		Status:    a.Status,
		Tokens:    a.Tokens,
		UserID:    a.UserID,
		ConvID:    a.ConvID,
		CreatedAt: a.CreatedAt,
		UpdatedAt: a.UpdatedAt,
		IsDeleted: a.IsDeleted,
	})
}

// UnmarshalJSON implements json.Unmarshaler interface for AiTask
func (a *AiTask) UnmarshalJSON(data []byte) error {
	var aiTask struct {
		ID        *int64     `json:"id"`
		Type      string     `json:"type"`
		Data      string     `json:"data"`
		Result    string     `json:"result"`
		Status    string     `json:"status"`
		Tokens    *int64     `json:"tokens"`
		UserID    int64      `json:"userId"`
		ConvID    int64      `json:"convId"`
		CreatedAt *time.Time `json:"createdAt"`
		UpdatedAt *time.Time `json:"updatedAt"`
		IsDeleted *bool      `json:"isDeleted"`
	}
	if err := json.Unmarshal(data, &aiTask); err != nil {
		return err
	}
	a.ID = aiTask.ID
	a.Type = aiTask.Type
	a.Data = aiTask.Data
	a.Result = aiTask.Result
	a.Status = aiTask.Status
	a.Tokens = aiTask.Tokens
	a.UserID = aiTask.UserID
	a.ConvID = aiTask.ConvID
	a.CreatedAt = aiTask.CreatedAt
	a.UpdatedAt = aiTask.UpdatedAt
	a.IsDeleted = aiTask.IsDeleted
	return nil
}
