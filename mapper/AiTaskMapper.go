package mapper

import (
	"drawsee/pojo/entity"
	"fmt"
	"time"
)

// AiTaskMapper 定义 AI 任务相关的数据操作接口
type AiTaskMapper interface {
	GetById(id int64) (*entity.AiTask, error)
	GetByUserIdAndConvIdAndStatus(userId, convId int64, status string) ([]*entity.AiTask, error)
	Insert(aiTask *entity.AiTask) error
	Update(aiTask *entity.AiTask) error
	CountTotalAiTasks() (int64, error)
	CountTaskTypeDistribution() ([]map[string]interface{}, error)
	CalculateAverageTaskDuration() (int64, error)
	SumTotalTokensConsumed() (int64, error)
	CountDailyTokenConsumption(days int) ([]map[string]interface{}, error)
	CountWeeklyTokenConsumption(weeks int) ([]map[string]interface{}, error)
	CountMonthlyTokenConsumption(months int) ([]map[string]interface{}, error)
	CountUsersWithAiTasks() (int64, error)
	CountDailySystemVisits(days int) ([]map[string]interface{}, error)
	CountWeeklySystemVisits(weeks int) ([]map[string]interface{}, error)
	CountMonthlySystemVisits(months int) ([]map[string]interface{}, error)
}

// MockAiTaskMapper 模拟 AiTaskMapper 实现，用于测试
type MockAiTaskMapper struct {
	tasks []*entity.AiTask
}

// NewMockAiTaskMapper 创建一个新的 MockAiTaskMapper 实例
func NewMockAiTaskMapper() *MockAiTaskMapper {
	return &MockAiTaskMapper{
		tasks: make([]*entity.AiTask, 0),
	}
}

// GetById 模拟根据 ID 获取 AI 任务
func (m *MockAiTaskMapper) GetById(id int64) (*entity.AiTask, error) {
	for _, task := range m.tasks {
		if task.ID != nil && *task.ID == id {
			return task, nil
		}
	}
	return nil, nil
}

// GetByUserIdAndConvIdAndStatus 模拟根据用户 ID、会话 ID 和状态获取 AI 任务列表
func (m *MockAiTaskMapper) GetByUserIdAndConvIdAndStatus(userId, convId int64, status string) ([]*entity.AiTask, error) {
	var result []*entity.AiTask
	for _, task := range m.tasks {
		if task.UserID == userId && task.ConvID == convId && task.Status == status {
			result = append(result, task)
		}
	}
	return result, nil
}

// Insert 模拟插入 AI 任务
func (m *MockAiTaskMapper) Insert(aiTask *entity.AiTask) error {
	m.tasks = append(m.tasks, aiTask)
	return nil
}

// Update 模拟更新 AI 任务
func (m *MockAiTaskMapper) Update(aiTask *entity.AiTask) error {
	for i, task := range m.tasks {
		if task.ID != nil && *task.ID == *aiTask.ID {
			m.tasks[i] = aiTask
			return nil
		}
	}
	return nil
}

// CountTotalAiTasks 模拟获取总 AI 任务数
func (m *MockAiTaskMapper) CountTotalAiTasks() (int64, error) {
	return int64(len(m.tasks)), nil
}

// CountTaskTypeDistribution 模拟获取不同类型任务的分布
func (m *MockAiTaskMapper) CountTaskTypeDistribution() ([]map[string]interface{}, error) {
	typeCount := make(map[string]int)
	for _, task := range m.tasks {
		typeCount[task.Type]++
	}
	var result []map[string]interface{}
	for taskType, count := range typeCount {
		result = append(result, map[string]interface{}{
			"type":  taskType,
			"count": count,
		})
	}
	return result, nil
}

// CalculateAverageTaskDuration 模拟获取任务平均耗时
func (m *MockAiTaskMapper) CalculateAverageTaskDuration() (int64, error) {
	var totalDuration int64
	var count int64
	for _, task := range m.tasks {
		if task.CreatedAt != nil && task.UpdatedAt != nil {
			duration := task.UpdatedAt.Sub(*task.CreatedAt).Milliseconds()
			totalDuration += duration
			count++
		}
	}
	if count == 0 {
		return 0, nil
	}
	return totalDuration / count, nil
}

// SumTotalTokensConsumed 模拟获取总 Token 消耗量
func (m *MockAiTaskMapper) SumTotalTokensConsumed() (int64, error) {
	var totalTokens int64
	for _, task := range m.tasks {
		if task.Tokens != nil {
			totalTokens += *task.Tokens
		}
	}
	return totalTokens, nil
}

// CountDailyTokenConsumption 模拟获取每日 Token 消耗统计
func (m *MockAiTaskMapper) CountDailyTokenConsumption(days int) ([]map[string]interface{}, error) {
	dailyConsumption := make(map[string]int64)
	now := time.Now()
	for _, task := range m.tasks {
		if task.CreatedAt != nil && task.Tokens != nil {
			day := task.CreatedAt.Format("2006-01-02")
			if now.Sub(*task.CreatedAt).Hours() <= 24*float64(days) {
				dailyConsumption[day] += *task.Tokens
			}
		}
	}
	var result []map[string]interface{}
	for day, consumption := range dailyConsumption {
		result = append(result, map[string]interface{}{
			"day":         day,
			"consumption": consumption,
		})
	}
	return result, nil
}

// CountWeeklyTokenConsumption 模拟获取每周 Token 消耗统计
func (m *MockAiTaskMapper) CountWeeklyTokenConsumption(weeks int) ([]map[string]interface{}, error) {
	weeklyConsumption := make(map[string]int64)
	now := time.Now()
	for _, task := range m.tasks {
		if task.CreatedAt != nil && task.Tokens != nil {
			year, week := task.CreatedAt.ISOWeek()
			weekKey := fmt.Sprintf("%d-W%d", year, week)
			if now.Sub(*task.CreatedAt).Hours() <= 24*7*float64(weeks) {
				weeklyConsumption[weekKey] += *task.Tokens
			}
		}
	}
	var result []map[string]interface{}
	for weekKey, consumption := range weeklyConsumption {
		result = append(result, map[string]interface{}{
			"week":        weekKey,
			"consumption": consumption,
		})
	}
	return result, nil
}

// CountMonthlyTokenConsumption 模拟获取每月 Token 消耗统计
func (m *MockAiTaskMapper) CountMonthlyTokenConsumption(months int) ([]map[string]interface{}, error) {
	monthlyConsumption := make(map[string]int64)
	now := time.Now()
	for _, task := range m.tasks {
		if task.CreatedAt != nil && task.Tokens != nil {
			monthKey := task.CreatedAt.Format("2006-01")
			if now.Sub(*task.CreatedAt).Hours() <= 24*30*float64(months) {
				monthlyConsumption[monthKey] += *task.Tokens
			}
		}
	}
	var result []map[string]interface{}
	for monthKey, consumption := range monthlyConsumption {
		result = append(result, map[string]interface{}{
			"month":       monthKey,
			"consumption": consumption,
		})
	}
	return result, nil
}

// CountUsersWithAiTasks 模拟获取使用过 AI 任务的用户数量
func (m *MockAiTaskMapper) CountUsersWithAiTasks() (int64, error) {
	userSet := make(map[int64]bool)
	for _, task := range m.tasks {
		userSet[task.UserID] = true
	}
	return int64(len(userSet)), nil
}

// CountDailySystemVisits 模拟获取系统每日访问量统计
func (m *MockAiTaskMapper) CountDailySystemVisits(days int) ([]map[string]interface{}, error) {
	dailyVisits := make(map[string]int)
	now := time.Now()
	for _, task := range m.tasks {
		if task.CreatedAt != nil {
			day := task.CreatedAt.Format("2006-01-02")
			if now.Sub(*task.CreatedAt).Hours() <= 24*float64(days) {
				dailyVisits[day]++
			}
		}
	}
	var result []map[string]interface{}
	for day, visits := range dailyVisits {
		result = append(result, map[string]interface{}{
			"day":    day,
			"visits": visits,
		})
	}
	return result, nil
}

// CountWeeklySystemVisits 模拟获取系统每周访问量统计
func (m *MockAiTaskMapper) CountWeeklySystemVisits(weeks int) ([]map[string]interface{}, error) {
	weeklyVisits := make(map[string]int)
	now := time.Now()
	for _, task := range m.tasks {
		if task.CreatedAt != nil {
			year, week := task.CreatedAt.ISOWeek()
			weekKey := fmt.Sprintf("%d-W%d", year, week)
			if now.Sub(*task.CreatedAt).Hours() <= 24*7*float64(weeks) {
				weeklyVisits[weekKey]++
			}
		}
	}
	var result []map[string]interface{}
	for weekKey, visits := range weeklyVisits {
		result = append(result, map[string]interface{}{
			"week":   weekKey,
			"visits": visits,
		})
	}
	return result, nil
}

// CountMonthlySystemVisits 模拟获取系统每月访问量统计
func (m *MockAiTaskMapper) CountMonthlySystemVisits(months int) ([]map[string]interface{}, error) {
	monthlyVisits := make(map[string]int)
	now := time.Now()
	for _, task := range m.tasks {
		if task.CreatedAt != nil {
			monthKey := task.CreatedAt.Format("2006-01")
			if now.Sub(*task.CreatedAt).Hours() <= 24*30*float64(months) {
				monthlyVisits[monthKey]++
			}
		}
	}
	var result []map[string]interface{}
	for monthKey, visits := range monthlyVisits {
		result = append(result, map[string]interface{}{
			"month":  monthKey,
			"visits": visits,
		})
	}
	return result, nil
}
