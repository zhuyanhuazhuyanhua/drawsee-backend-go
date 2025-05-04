package entity

import (
	"time"
)

// InvitationCode 表示邀请码实体
type InvitationCode struct {
	ID           *int64     `json:"id"`
	Code         string     `json:"code"`
	CreatedAt    *time.Time `json:"created_at"`
	UsedBy       *int64     `json:"used_by"`
	UsedAt       *time.Time `json:"used_at"`
	IsActive     *bool      `json:"is_active"`
	SentUserName string     `json:"sent_user_name"`
	SentEmail    string     `json:"sent_email"`
	LastSentAt   *time.Time `json:"last_sent_at"`
}

// NewInvitationCode 创建一个新的 InvitationCode 实例
func NewInvitationCode(code string) *InvitationCode {
	now := time.Now()
	isActive := true
	return &InvitationCode{
		Code:      code,
		CreatedAt: &now,
		IsActive:  &isActive,
	}
}
