package entity_test

import (
	"drawsee/pojo/entity"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	username := "testUser"
	password := "testPassword"

	user := entity.NewUser(username, password)

	assert.Equal(t, username, user.Username)
	assert.Equal(t, password, user.Password)
	assert.NotNil(t, user.CreatedAt)
	assert.NotNil(t, user.UpdatedAt)
	assert.False(t, *user.IsDeleted)

	// 检查创建时间和更新时间是否相近
	diff := user.UpdatedAt.Sub(*user.CreatedAt)
	assert.True(t, diff < time.Second, "创建时间和更新时间应该相近")
}
