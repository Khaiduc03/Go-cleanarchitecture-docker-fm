package userImpl

import (
	"FM/src/entities"
	"FM/src/modules/user"
	modelUser "FM/src/modules/user/model"
	"context"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	*gorm.DB
}


func (userRepository *UserRepositoryImpl) GetProfile(ctx context.Context, id uint) (user entities.User, err error) {
	result := userRepository.DB.WithContext(ctx).Where("id = ?", id).First(&user)
	if result.RowsAffected == 0 {
		return entities.User{}, nil
	}
	return user, nil
}


func (userRepository *UserRepositoryImpl) UpdateProfile(ctx context.Context, id uint, req modelUser.UpdateUserReq) (string, error) {
	var user entities.User
	result := userRepository.DB.WithContext(ctx).Where("id = ?", id).First(&user)
	if result.RowsAffected == 0 {
		return "", nil
	}
	userRepository.DB.WithContext(ctx).Model(&user).Updates(req)
	return "Update profile successfully", nil
}

func NewUserRepositoryImpl(db *gorm.DB) user.UserRepository {
	return &UserRepositoryImpl{db}
}

func (userRepository *UserRepositoryImpl) GetAllStaff(ctx context.Context) ([]entities.User, error) {
	var users []entities.User
	err := userRepository.DB.WithContext(ctx).Where("role = 'STAFF'").Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

