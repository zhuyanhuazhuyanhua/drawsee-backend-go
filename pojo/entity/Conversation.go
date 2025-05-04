package entity

import (
	"time"
)

// Conversation 表示会话实体
type Conversation struct {
	ID        *int64     `json:"id"`
	Title     string     `json:"title"`
	UserID    *int64     `json:"user_id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	IsDeleted *bool      `json:"is_deleted"`
}

// NewConversation 创建一个新的 Conversation 实例
func NewConversation(title string, userID int64) *Conversation {
	now := time.Now()
	isDeleted := false
	return &Conversation{
		Title:     title,
		UserID:    &userID,
		CreatedAt: &now,
		UpdatedAt: &now,
		IsDeleted: &isDeleted,
	}
}
