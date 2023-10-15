package AuthImpl

import (
	"context"

	Auth "FM/src/auth"
	"FM/src/auth/models"
	"FM/src/entities"

	firebase "firebase.google.com/go/v4"

	"gorm.io/gorm"
)

type AuthRepositoryImpl struct {
	*gorm.DB
	*firebase.App
}

func NewAuthRepositoryImpl(DB *gorm.DB) Auth.AuthRepository {
	return &AuthRepositoryImpl{
		DB: DB,
	}
}

func (authRepository *AuthRepositoryImpl) SignInWithGoogle(ctx context.Context, model models.Payload) (entities.User, error) {
	var user entities.User

	email := model.Email
	isExist := authRepository.DB.WithContext(ctx).Where("email = ?", email).Find(&user)

	if isExist.Error != nil {
		return entities.User{}, isExist.Error
	}

	if isExist.RowsAffected == 0 {
		newUser := entities.User{
			Email: email,
			Url:   model.Picture,
			Name:  model.Name,
		}

		result := authRepository.DB.WithContext(ctx).Create(&newUser)
		if result.Error != nil {
			return entities.User{}, result.Error
		}
		user = newUser
	} else {
		user.Url = model.Picture
		user.Name = model.Name
		result := authRepository.DB.WithContext(ctx).Save(&user)
		if result.Error != nil {
			return entities.User{}, result.Error
		}
	}
	return user, nil
}
