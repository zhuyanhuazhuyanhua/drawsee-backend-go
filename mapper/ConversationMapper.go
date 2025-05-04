package mapper

import (
	"drawsee/pojo/entity"
	"fmt"
	"time"
)

// ConversationMapper 定义会话相关的数据操作接口
type ConversationMapper interface {
	GetById(id int64) (*entity.Conversation, error)
	GetByUserId(userId int64) ([]*entity.Conversation, error)
	Insert(conversation *entity.Conversation) error
	Update(conversation *entity.Conversation) error
	CountTotalConversations() (int64, error)
	CountNewConversationsBetween(startTime, endTime time.Time) (int64, error)
	CountDailyNewConversations(days int) ([]map[string]interface{}, error)
	CountWeeklyNewConversations(weeks int) ([]map[string]interface{}, error)
	CountMonthlyNewConversations(months int) ([]map[string]interface{}, error)
	CountUsersWithConversations() (int64, error)
}

// MockConversationMapper 模拟 ConversationMapper 实现，用于测试
type MockConversationMapper struct {
	conversations []*entity.Conversation
}

// NewMockConversationMapper 创建一个新的 MockConversationMapper 实例
func NewMockConversationMapper() *MockConversationMapper {
	return &MockConversationMapper{
		conversations: make([]*entity.Conversation, 0),
	}
}

// GetById 模拟根据 ID 获取会话
func (m *MockConversationMapper) GetById(id int64) (*entity.Conversation, error) {
	for _, conv := range m.conversations {
		if conv.ID != nil && *conv.ID == id {
			return conv, nil
		}
	}
	return nil, nil
}

// GetByUserId 模拟根据用户 ID 获取会话列表
func (m *MockConversationMapper) GetByUserId(userId int64) ([]*entity.Conversation, error) {
	var result []*entity.Conversation
	for _, conv := range m.conversations {
		if conv.UserID != nil && *conv.UserID == userId {
			result = append(result, conv)
		}
	}
	return result, nil
}

// Insert 模拟插入会话
func (m *MockConversationMapper) Insert(conversation *entity.Conversation) error {
	m.conversations = append(m.conversations, conversation)
	return nil
}

// Update 模拟更新会话
func (m *MockConversationMapper) Update(conversation *entity.Conversation) error {
	for i, conv := range m.conversations {
		if conv.ID != nil && *conv.ID == *conversation.ID {
			m.conversations[i] = conversation
			return nil
		}
	}
	return nil
}

// CountTotalConversations 模拟获取总会话数
func (m *MockConversationMapper) CountTotalConversations() (int64, error) {
	return int64(len(m.conversations)), nil
}

// CountNewConversationsBetween 模拟获取某个时间段内的新增会话数
func (m *MockConversationMapper) CountNewConversationsBetween(startTime, endTime time.Time) (int64, error) {
	var count int64
	for _, conv := range m.conversations {
		if conv.CreatedAt != nil && !conv.CreatedAt.Before(startTime) && conv.CreatedAt.Before(endTime) {
			count++
		}
	}
	return count, nil
}

// CountDailyNewConversations 模拟获取每日新增会话统计
func (m *MockConversationMapper) CountDailyNewConversations(days int) ([]map[string]interface{}, error) {
	dailyCount := make(map[string]int64)
	now := time.Now()
	startDate := now.AddDate(0, 0, -days)

	for _, conv := range m.conversations {
		if conv.CreatedAt != nil && !conv.CreatedAt.Before(startDate) {
			day := conv.CreatedAt.Format("2006-01-02")
			dailyCount[day]++
		}
	}

	var result []map[string]interface{}
	for day, count := range dailyCount {
		result = append(result, map[string]interface{}{
			"day":   day,
			"count": count,
		})
	}
	return result, nil
}

// CountWeeklyNewConversations 模拟获取每周新增会话统计
func (m *MockConversationMapper) CountWeeklyNewConversations(weeks int) ([]map[string]interface{}, error) {
	weeklyCount := make(map[string]int64)
	now := time.Now()
	startDate := now.AddDate(0, 0, -weeks*7)

	for _, conv := range m.conversations {
		if conv.CreatedAt != nil && !conv.CreatedAt.Before(startDate) {
			year, week := conv.CreatedAt.ISOWeek()
			weekKey := fmt.Sprintf("%d-W%d", year, week)
			weeklyCount[weekKey]++
		}
	}

	var result []map[string]interface{}
	for weekKey, count := range weeklyCount {
		result = append(result, map[string]interface{}{
			"week":  weekKey,
			"count": count,
		})
	}
	return result, nil
}

// CountMonthlyNewConversations 模拟获取每月新增会话统计
func (m *MockConversationMapper) CountMonthlyNewConversations(months int) ([]map[string]interface{}, error) {
	monthlyCount := make(map[string]int64)
	now := time.Now()
	startDate := now.AddDate(0, -months, 0)

	for _, conv := range m.conversations {
		if conv.CreatedAt != nil && !conv.CreatedAt.Before(startDate) {
			monthKey := conv.CreatedAt.Format("2006-01")
			monthlyCount[monthKey]++
		}
	}

	var result []map[string]interface{}
	for monthKey, count := range monthlyCount {
		result = append(result, map[string]interface{}{
			// Bug 修复：使用 monthKey 作为 "month" 键的值
			"month": monthKey,
			"count": count,
		})
	}
	return result, nil
}

// CountUsersWithConversations 模拟获取有会话的用户数量
func (m *MockConversationMapper) CountUsersWithConversations() (int64, error) {
	userSet := make(map[int64]bool)
	for _, conv := range m.conversations {
		if conv.UserID != nil {
			userSet[*conv.UserID] = true
		}
	}
	return int64(len(userSet)), nil
}
