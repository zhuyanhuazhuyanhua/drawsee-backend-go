package repository

import (
	"drawsee/internal/app/model"

	"gorm.io/gorm"
)

// UserRepository 用户数据访问接口
type UserRepository interface {
	CreateUser(user *model.User) error
	GetUserByUsername(username string) (*model.User, error)
	// 可根据需求添加更多方法
}

// UserRepoImpl 用户数据访问实现
type UserRepoImpl struct {
	db *gorm.DB
}

// NewUserRepoImpl 创建用户数据访问实例
func NewUserRepoImpl(db *gorm.DB) *UserRepoImpl {
	return &UserRepoImpl{db: db}
}

// CreateUser 创建用户
func (r *UserRepoImpl) CreateUser(user *model.User) error {
	return r.db.Create(user).Error
}

// GetUserByUsername 根据用户名获取用户
func (r *UserRepoImpl) GetUserByUsername(username string) (*model.User, error) {
	var user model.User
	err := r.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
