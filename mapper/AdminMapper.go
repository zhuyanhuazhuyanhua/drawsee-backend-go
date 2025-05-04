package mapper

import (
	"drawsee/pojo/entity"
	"errors"
)

// AdminMapper 定义 Admin 相关的数据操作接口
type AdminMapper interface {
	// GetByUserId 根据用户 ID 获取 Admin 实体
	GetByUserId(userId int64) (*entity.Admin, error)
	// Insert 插入一个 Admin 实体
	Insert(admin *entity.Admin) error
}

// MockAdminMapper 模拟 AdminMapper 实现，用于测试
type MockAdminMapper struct {
	admins map[int64]*entity.Admin
}

// NewMockAdminMapper 创建一个新的 MockAdminMapper 实例
func NewMockAdminMapper() *MockAdminMapper {
	return &MockAdminMapper{
		admins: make(map[int64]*entity.Admin),
	}
}

// GetByUserId 模拟根据用户 ID 获取 Admin 实体
func (m *MockAdminMapper) GetByUserId(userId int64) (*entity.Admin, error) {
	admin, exists := m.admins[userId]
	if !exists {
		return nil, errors.New("admin not found")
	}
	return admin, nil
}

// Insert 模拟插入一个 Admin 实体
func (m *MockAdminMapper) Insert(admin *entity.Admin) error {
	if admin == nil {
		return errors.New("admin cannot be nil")
	}
	m.admins[admin.UserID] = admin
	return nil
}

// AdminExists 检查指定用户 ID 的 Admin 是否存在
func (m *MockAdminMapper) AdminExists(userId int64) bool {
	_, exists := m.admins[userId]
	return exists
}
