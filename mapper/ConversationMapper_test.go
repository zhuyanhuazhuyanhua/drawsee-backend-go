package mapper_test

import (
	"drawsee/mapper"
	"drawsee/pojo/entity"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMockConversationMapper(t *testing.T) {
	mockMapper := mapper.NewMockConversationMapper()

	// 准备测试数据
	id := int64(1)
	title := "Test Conversation"
	userId := int64(100)
	createdAt := time.Now()
	updatedAt := createdAt
	isDeleted := false

	conversation := &entity.Conversation{
		ID:        &id,
		Title:     title,
		UserID:    &userId,
		CreatedAt: &createdAt,
		UpdatedAt: &updatedAt,
		IsDeleted: &isDeleted,
	}

	// 测试 Insert
	err := mockMapper.Insert(conversation)
	assert.NoError(t, err)

	// 测试 GetById
	foundConv, err := mockMapper.GetById(id)
	assert.NoError(t, err)
	assert.Equal(t, conversation, foundConv)

	// 测试 GetByUserId
	convs, err := mockMapper.GetByUserId(userId)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(convs))

	// 测试 Update
	newTitle := "Updated Conversation"
	conversation.Title = newTitle
	err = mockMapper.Update(conversation)
	assert.NoError(t, err)
	updatedConv, err := mockMapper.GetById(id)
	assert.NoError(t, err)
	assert.Equal(t, newTitle, updatedConv.Title)

	// 测试 CountTotalConversations
	totalConvs, err := mockMapper.CountTotalConversations()
	assert.NoError(t, err)
	assert.Equal(t, int64(1), totalConvs)

	// 测试 CountNewConversationsBetween
	startTime := createdAt.Add(-1 * time.Hour)
	endTime := createdAt.Add(1 * time.Hour)
	newConvsCount, err := mockMapper.CountNewConversationsBetween(startTime, endTime)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), newConvsCount)

	// 测试 CountDailyNewConversations
	dailyConvs, err := mockMapper.CountDailyNewConversations(1)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(dailyConvs))

	// 测试 CountWeeklyNewConversations
	weeklyConvs, err := mockMapper.CountWeeklyNewConversations(1)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(weeklyConvs))

	// 测试 CountMonthlyNewConversations
	monthlyConvs, err := mockMapper.CountMonthlyNewConversations(1)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(monthlyConvs))

	// 测试 CountUsersWithConversations
	userCount, err := mockMapper.CountUsersWithConversations()
	assert.NoError(t, err)
	assert.Equal(t, int64(1), userCount)
}
