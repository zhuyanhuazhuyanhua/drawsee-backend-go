package model

import (
	"time"
)

// User 用户表结构
type User struct {
	ID       uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Username string `gorm:"unique;size:50" json:"username"`
	Password string `gorm:"size:255" json:"password"`
	// 这里需要补充完整 `user` 表中未给出的字段
}

// Admin 管理员表结构
type Admin struct {
	ID     int  `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID uint `gorm:"unique" json:"user_id"`
}

// AiTask AI 任务表结构
type AiTask struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Type      string    `gorm:"size:50" json:"type"`
	Data      string    `gorm:"type:mediumtext" json:"data"`
	Result    *string   `gorm:"type:mediumtext" json:"result"`
	Status    string    `gorm:"size:50" json:"status"`
	Tokens    *uint     `json:"tokens"`
	UserID    uint      `json:"user_id"`
	ConvID    uint      `json:"conv_id"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP;autoUpdateTime" json:"updated_at"`
	IsDeleted bool      `gorm:"default:false" json:"is_deleted"`
}

// Conversation 会话表结构
type Conversation struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Title     string    `gorm:"size:255" json:"title"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP;autoUpdateTime" json:"updated_at"`
	IsDeleted bool      `gorm:"default:false" json:"is_deleted"`
}

// InvitationCode 邀请码表结构
type InvitationCode struct {
	ID           int        `gorm:"primaryKey;autoIncrement" json:"id"`
	Code         string     `gorm:"unique;size:255" json:"code"`
	CreatedAt    time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UsedBy       *int       `json:"used_by"`
	UsedAt       *time.Time `json:"used_at"`
	IsActive     bool       `gorm:"default:true" json:"is_active"`
	SentUserName *string    `gorm:"size:255" json:"sent_user_name"`
	SentEmail    *string    `gorm:"size:255" json:"sent_email"`
	LastSentAt   *time.Time `json:"last_sent_at"`
}

// Node 节点表结构
type Node struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Type      string    `gorm:"size:50" json:"type"`
	Data      string    `gorm:"type:mediumtext" json:"data"`
	Position  string    `gorm:"type:text" json:"position"`
	Height    *uint     `json:"height"`
	ParentID  *uint     `json:"parent_id"`
	ConvID    uint      `json:"conv_id"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP;autoUpdateTime" json:"updated_at"`
	IsDeleted bool      `gorm:"default:false" json:"is_deleted"`
}
