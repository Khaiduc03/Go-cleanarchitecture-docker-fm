package user

import (
	"FM/src/entities"
	modelUser "FM/src/modules/user/model"
	"context"
)

type UserRepository interface {
	GetProfile(ctx context.Context, id uint) (user entities.User,err error)
	UpdateProfile(ctx context.Context, id uint, req modelUser.UpdateUserReq) (string, error)
	GetAllStaff(ctx context.Context) ([]entities.User, error)
}
