package userImpl

import (
	"FM/src/entities"
	"FM/src/modules/user"
	modelUser "FM/src/modules/user/model"
	"context"
)

type UserServiceImpl struct {
	user.UserRepository
}

func NewUserServiceImpl(userRepository *user.UserRepository) user.UserService {
	return &UserServiceImpl{UserRepository: *userRepository}
}

func (userService *UserServiceImpl) GetProfile(ctx context.Context, id uint) (entities.User, error) {
	return userService.UserRepository.GetProfile(ctx, id)
}

func (userService *UserServiceImpl) UpdateProfile(ctx context.Context, id uint, req modelUser.UpdateUserReq) (string, error) {
	return userService.UserRepository.UpdateProfile(ctx, id, req)
}

func (userService *UserServiceImpl) GetAllStaff(ctx context.Context) ([]entities.User, error) {
	return userService.UserRepository.GetAllStaff(ctx)
}