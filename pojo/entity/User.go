package entity

import (
	"time"
)

// User 表示用户实体
type User struct {
	ID        *int64     `json:"id"`
	Username  string     `json:"username"`
	Password  string     `json:"password"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	IsDeleted *bool      `json:"is_deleted"`
}

// NewUser 创建一个新的 User 实例
func NewUser(username, password string) *User {
	now := time.Now()
	isDeleted := false
	return &User{
		Username:  username,
		Password:  password,
		CreatedAt: &now,
		UpdatedAt: &now,
		IsDeleted: &isDeleted,
	}
}
