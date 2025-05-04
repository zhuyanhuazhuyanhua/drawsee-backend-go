package mapper_test

import (
	"drawsee/mapper"
	"drawsee/pojo/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMockAdminMapper_GetByUserId(t *testing.T) {
	mockMapper := mapper.NewMockAdminMapper()
	userId := int64(123)
	admin := entity.NewAdmin(userId)

	err := mockMapper.Insert(admin)
	assert.NoError(t, err)

	foundAdmin, err := mockMapper.GetByUserId(userId)
	assert.NoError(t, err)
	assert.Equal(t, admin, foundAdmin)

	nonExistentUserId := int64(456)
	_, err = mockMapper.GetByUserId(nonExistentUserId)
	assert.Error(t, err)
}

func TestMockAdminMapper_Insert(t *testing.T) {
	mockMapper := mapper.NewMockAdminMapper()
	userId := int64(123)
	admin := entity.NewAdmin(userId)

	err := mockMapper.Insert(admin)
	assert.NoError(t, err)

	// 使用 AdminExists 方法检查 Admin 是否存在
	exists := mockMapper.AdminExists(userId)
	assert.True(t, exists)

	err = mockMapper.Insert(nil)
	assert.Error(t, err)
}
