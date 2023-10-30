package user

import (
	"FM/src/entities"
	modelUser "FM/src/modules/user/model"
	"context"
)

type UserService interface {
	GetProfile(ctx context.Context, id uint) (entities.User, error)
	UpdateProfile(ctx context.Context, id uint, req modelUser.UpdateUserReq) (string, error)
	GetAllStaff(ctx context.Context) ([]entities.User, error)
}
