package entity_test

import (
	"drawsee/pojo/entity"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewConversation(t *testing.T) {
	title := "Test Conversation"
	userID := int64(123)

	conv := entity.NewConversation(title, userID)

	assert.Equal(t, title, conv.Title)
	assert.Equal(t, userID, *conv.UserID)
	assert.NotNil(t, conv.CreatedAt)
	assert.NotNil(t, conv.UpdatedAt)
	assert.False(t, *conv.IsDeleted)

	// 检查创建时间和更新时间是否相近
	diff := conv.UpdatedAt.Sub(*conv.CreatedAt)
	assert.True(t, diff < time.Second, "CreatedAt and UpdatedAt should be close")
}
