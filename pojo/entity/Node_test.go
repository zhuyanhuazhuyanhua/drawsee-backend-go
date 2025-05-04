package entity_test

import (
	"drawsee/pojo/entity"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewNode(t *testing.T) {
	nodeType := "test_type"
	data := "test_data"
	position := "test_position"
	parentID := int64(1)
	userID := int64(2)
	convID := int64(3)
	isDeleted := false

	node := entity.NewNode(nodeType, data, position, parentID, userID, convID, isDeleted)

	assert.Equal(t, nodeType, node.Type)
	assert.Equal(t, data, node.Data)
	assert.Equal(t, position, node.Position)
	assert.Equal(t, parentID, *node.ParentID)
	assert.Equal(t, userID, *node.UserID)
	assert.Equal(t, convID, *node.ConvID)
	assert.Equal(t, isDeleted, *node.IsDeleted)
	assert.NotNil(t, node.CreatedAt)
	assert.NotNil(t, node.UpdatedAt)

	// 检查创建时间和更新时间是否相近
	diff := node.UpdatedAt.Sub(*node.CreatedAt)
	assert.True(t, diff < time.Second, "CreatedAt and UpdatedAt should be close")
}
