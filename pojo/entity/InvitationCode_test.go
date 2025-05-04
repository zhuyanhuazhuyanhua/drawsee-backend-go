package entity_test

import (
	"drawsee/pojo/entity"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewInvitationCode(t *testing.T) {
	code := "TEST123"
	invitationCode := entity.NewInvitationCode(code)

	assert.Equal(t, code, invitationCode.Code)
	assert.NotNil(t, invitationCode.CreatedAt)
	assert.True(t, *invitationCode.IsActive)

	// 检查创建时间是否合理
	assert.True(t, time.Since(*invitationCode.CreatedAt) < time.Second, "创建时间应该接近当前时间")
}
