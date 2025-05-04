package mapper_test

import (
	"drawsee/mapper"
	"drawsee/pojo/entity"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMockAiTaskMapper(t *testing.T) {
	mockMapper := mapper.NewMockAiTaskMapper()

	// 准备测试数据
	id := int64(1)
	userId := int64(100)
	convId := int64(200)
	status := "completed"
	taskType := "test_type"
	data := "test_data"
	result := "test_result"
	tokens := int64(1000)
	createdAt := time.Now()
	updatedAt := createdAt.Add(1 * time.Hour)
	isDeleted := false

	aiTask := &entity.AiTask{
		ID:        &id,
		Type:      taskType,
		Data:      data,
		Result:    result,
		Status:    status,
		Tokens:    &tokens,
		UserID:    userId,
		ConvID:    convId,
		CreatedAt: &createdAt,
		UpdatedAt: &updatedAt,
		IsDeleted: &isDeleted,
	}

	// 测试 Insert
	err := mockMapper.Insert(aiTask)
	assert.NoError(t, err)

	// 测试 GetById
	foundTask, err := mockMapper.GetById(id)
	assert.NoError(t, err)
	assert.Equal(t, aiTask, foundTask)

	// 测试 GetByUserIdAndConvIdAndStatus
	tasks, err := mockMapper.GetByUserIdAndConvIdAndStatus(userId, convId, status)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(tasks))

	// 测试 Update
	newStatus := "updated"
	aiTask.Status = newStatus
	err = mockMapper.Update(aiTask)
	assert.NoError(t, err)
	updatedTask, err := mockMapper.GetById(id)
	assert.NoError(t, err)
	assert.Equal(t, newStatus, updatedTask.Status)

	// 测试 CountTotalAiTasks
	totalTasks, err := mockMapper.CountTotalAiTasks()
	assert.NoError(t, err)
	assert.Equal(t, int64(1), totalTasks)

	// 测试 CountTaskTypeDistribution
	taskTypeDist, err := mockMapper.CountTaskTypeDistribution()
	assert.NoError(t, err)
	assert.Equal(t, 1, len(taskTypeDist))

	// 测试 CalculateAverageTaskDuration
	avgDuration, err := mockMapper.CalculateAverageTaskDuration()
	assert.NoError(t, err)
	assert.Equal(t, int64(3600000), avgDuration)

	// 测试 SumTotalTokensConsumed
	totalTokens, err := mockMapper.SumTotalTokensConsumed()
	assert.NoError(t, err)
	assert.Equal(t, int64(1000), totalTokens)

	// 测试 CountDailyTokenConsumption
	dailyTokenConsumption, err := mockMapper.CountDailyTokenConsumption(1)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(dailyTokenConsumption))

	// 测试 CountWeeklyTokenConsumption
	weeklyTokenConsumption, err := mockMapper.CountWeeklyTokenConsumption(1)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(weeklyTokenConsumption))

	// 测试 CountMonthlyTokenConsumption
	monthlyTokenConsumption, err := mockMapper.CountMonthlyTokenConsumption(1)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(monthlyTokenConsumption))

	// 测试 CountUsersWithAiTasks
	userCount, err := mockMapper.CountUsersWithAiTasks()
	assert.NoError(t, err)
	assert.Equal(t, int64(1), userCount)

	// 测试 CountDailySystemVisits
	dailyVisits, err := mockMapper.CountDailySystemVisits(1)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(dailyVisits))

	// 测试 CountWeeklySystemVisits
	weeklyVisits, err := mockMapper.CountWeeklySystemVisits(1)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(weeklyVisits))

	// 测试 CountMonthlySystemVisits
	monthlyVisits, err := mockMapper.CountMonthlySystemVisits(1)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(monthlyVisits))
}
