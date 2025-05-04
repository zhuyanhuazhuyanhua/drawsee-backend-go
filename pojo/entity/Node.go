package entity

import (
	"time"
)

// Node 表示节点实体
type Node struct {
	ID        *int64     `json:"id"`
	Type      string     `json:"type"`
	Data      string     `json:"data"`
	Position  string     `json:"position"`
	Height    *int64     `json:"height"`
	Width     *int64     `json:"width"`
	ParentID  *int64     `json:"parent_id"`
	ConvID    *int64     `json:"conv_id"`
	UserID    *int64     `json:"user_id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	IsDeleted *bool      `json:"is_deleted"`
}

// NewNode 创建一个新的 Node 实例
func NewNode(
	nodeType, data, position string,
	parentID, userID, convID int64,
	isDeleted bool,
) *Node {
	now := time.Now()
	return &Node{
		Type:      nodeType,
		Data:      data,
		Position:  position,
		ParentID:  &parentID,
		UserID:    &userID,
		ConvID:    &convID,
		IsDeleted: &isDeleted,
		CreatedAt: &now,
		UpdatedAt: &now,
	}
}
