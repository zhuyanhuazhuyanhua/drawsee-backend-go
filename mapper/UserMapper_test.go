package mapper_test

import (
	"drawsee/mapper"
	"drawsee/pojo/entity"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMockUserMapper(t *testing.T) {
	mockMapper := mapper.NewMockUserMapper()

	// Prepare test data
	id := int64(1)
	username := "testUser"
	password := "testPassword"
	createdAt := time.Now()
	// Set UpdatedAt to a time after baseDate + afterDays to ensure the user is active
	afterDays := 1
	updatedAt := createdAt.AddDate(0, 0, afterDays+1)
	isDeleted := false

	user := &entity.User{
		ID:        &id,
		Username:  username,
		Password:  password,
		CreatedAt: &createdAt,
		UpdatedAt: &updatedAt,
		IsDeleted: &isDeleted,
	}

	// Test Insert
	err := mockMapper.Insert(user)
	assert.NoError(t, err)

	// Test GetById
	foundUser, err := mockMapper.GetById(id)
	assert.NoError(t, err)
	assert.Equal(t, user, foundUser)

	// Test GetByUsername
	foundByUsername, err := mockMapper.GetByUsername(username)
	assert.NoError(t, err)
	assert.Equal(t, user, foundByUsername)

	// Test CountTotalUsers
	totalUsers, err := mockMapper.CountTotalUsers()
	assert.NoError(t, err)
	assert.Equal(t, int64(1), totalUsers)

	// Test CountNewUsersBetween
	startTime := createdAt.Add(-1 * time.Hour)
	endTime := createdAt.Add(1 * time.Hour)
	newUsersCount, err := mockMapper.CountNewUsersBetween(startTime, endTime)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), newUsersCount)

	// Test CountDailyNewUsers
	dailyUsers, err := mockMapper.CountDailyNewUsers(1)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(dailyUsers))

	// Test CountWeeklyNewUsers
	weeklyUsers, err := mockMapper.CountWeeklyNewUsers(1)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(weeklyUsers))

	// Test CountMonthlyNewUsers
	monthlyUsers, err := mockMapper.CountMonthlyNewUsers(1)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(monthlyUsers))

	// Test CountActiveUsersBetween
	activeUsers, err := mockMapper.CountActiveUsersBetween(startTime, endTime)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), activeUsers)

	// Test CalculateRetentionRate
	baseDate := createdAt
	retentionRate, err := mockMapper.CalculateRetentionRate(baseDate, afterDays)
	assert.NoError(t, err)
	assert.Equal(t, 1.0, retentionRate)
}
