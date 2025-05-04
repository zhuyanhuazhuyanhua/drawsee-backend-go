package mapper

import (
	"drawsee/pojo/entity"
	"fmt"
	"time"
)

// UserMapper 定义用户相关的数据操作接口
type UserMapper interface {
	GetById(id int64) (*entity.User, error)
	GetByUsername(username string) (*entity.User, error)
	Insert(user *entity.User) error
	CountTotalUsers() (int64, error)
	CountNewUsersBetween(startTime, endTime time.Time) (int64, error)
	CountDailyNewUsers(days int) ([]map[string]interface{}, error)
	CountWeeklyNewUsers(weeks int) ([]map[string]interface{}, error)
	CountMonthlyNewUsers(months int) ([]map[string]interface{}, error)
	CountActiveUsersBetween(startTime, endTime time.Time) (int64, error)
	CalculateRetentionRate(baseDate time.Time, afterDays int) (float64, error)
}

// MockUserMapper 模拟 UserMapper 实现，用于测试
type MockUserMapper struct {
	users       map[int64]*entity.User
	usernameMap map[string]*entity.User
}

// NewMockUserMapper 创建一个新的 MockUserMapper 实例
func NewMockUserMapper() *MockUserMapper {
	return &MockUserMapper{
		users:       make(map[int64]*entity.User),
		usernameMap: make(map[string]*entity.User),
	}
}

// GetById 模拟根据 ID 获取用户
func (m *MockUserMapper) GetById(id int64) (*entity.User, error) {
	user, exists := m.users[id]
	if !exists {
		return nil, nil
	}
	return user, nil
}

// GetByUsername 模拟根据用户名获取用户
func (m *MockUserMapper) GetByUsername(username string) (*entity.User, error) {
	user, exists := m.usernameMap[username]
	if !exists {
		return nil, nil
	}
	return user, nil
}

// Insert 模拟插入用户
func (m *MockUserMapper) Insert(user *entity.User) error {
	if user.ID != nil {
		m.users[*user.ID] = user
	}
	if user.Username != "" {
		m.usernameMap[user.Username] = user
	}
	return nil
}

// CountTotalUsers 模拟获取总用户数
func (m *MockUserMapper) CountTotalUsers() (int64, error) {
	return int64(len(m.users)), nil
}

// CountNewUsersBetween 模拟获取某个时间段内的新增用户数
func (m *MockUserMapper) CountNewUsersBetween(startTime, endTime time.Time) (int64, error) {
	var count int64
	for _, user := range m.users {
		if user.CreatedAt != nil && !user.CreatedAt.Before(startTime) && user.CreatedAt.Before(endTime) {
			count++
		}
	}
	return count, nil
}

// CountDailyNewUsers 模拟获取每日新增用户统计
func (m *MockUserMapper) CountDailyNewUsers(days int) ([]map[string]interface{}, error) {
	dailyCount := make(map[string]int64)
	now := time.Now()
	startDate := now.AddDate(0, 0, -days)

	for _, user := range m.users {
		if user.CreatedAt != nil && !user.CreatedAt.Before(startDate) {
			day := user.CreatedAt.Format("2006-01-02")
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

// CountWeeklyNewUsers 模拟获取每周新增用户统计
func (m *MockUserMapper) CountWeeklyNewUsers(weeks int) ([]map[string]interface{}, error) {
	weeklyCount := make(map[string]int64)
	now := time.Now()
	startDate := now.AddDate(0, 0, -weeks*7)

	for _, user := range m.users {
		if user.CreatedAt != nil && !user.CreatedAt.Before(startDate) {
			year, week := user.CreatedAt.ISOWeek()
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

// CountMonthlyNewUsers 模拟获取每月新增用户统计
func (m *MockUserMapper) CountMonthlyNewUsers(months int) ([]map[string]interface{}, error) {
	monthlyCount := make(map[string]int64)
	now := time.Now()
	startDate := now.AddDate(0, -months, 0)

	for _, user := range m.users {
		if user.CreatedAt != nil && !user.CreatedAt.Before(startDate) {
			monthKey := user.CreatedAt.Format("2006-01")
			monthlyCount[monthKey]++
		}
	}

	var result []map[string]interface{}
	for monthKey, count := range monthlyCount {
		result = append(result, map[string]interface{}{
			// Bug fix: Use monthKey as the value for the "month" key
			"month": monthKey,
			"count": count,
		})
	}
	return result, nil
}

// CountActiveUsersBetween 模拟获取特定时间段内活跃的用户数(有 AI 任务记录)
// 这里简单假设所有用户都活跃，实际需要根据 AI 任务记录实现
func (m *MockUserMapper) CountActiveUsersBetween(startTime, endTime time.Time) (int64, error) {
	var count int64
	for _, user := range m.users {
		if user.CreatedAt != nil && !user.CreatedAt.Before(startTime) && user.CreatedAt.Before(endTime) {
			count++
		}
	}
	return count, nil
}

// CalculateRetentionRate 模拟获取留存率计算所需数据
func (m *MockUserMapper) CalculateRetentionRate(baseDate time.Time, afterDays int) (float64, error) {
	registeredUsers := 0
	activeUsers := 0
	afterDate := baseDate.AddDate(0, 0, afterDays)

	for _, user := range m.users {
		if user.CreatedAt != nil && user.CreatedAt.Format("2006-01-02") == baseDate.Format("2006-01-02") {
			registeredUsers++
			// 简单假设更新时间在 afterDate 之后的用户为活跃用户
			if user.UpdatedAt != nil && !user.UpdatedAt.Before(afterDate) {
				activeUsers++
			}
		}
	}

	if registeredUsers == 0 {
		return 0, nil
	}
	return float64(activeUsers) / float64(registeredUsers), nil
}
