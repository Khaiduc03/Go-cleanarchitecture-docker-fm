package AuthImpl

import (
	Auth "FM/src/auth"
	"FM/src/auth/models"
	"FM/src/entities"
	"context"
	"errors"
)

type AuthServiceImpl struct {
	Auth.AuthRepository
}

func NewAuthServiceImpl(authRepository *Auth.AuthRepository) Auth.AuthService {
	return &AuthServiceImpl{AuthRepository: *authRepository}
}

func (authService *AuthServiceImpl) SignInWithGoogle(ctx context.Context, model models.SignInWithGoogleModel) (entities.User, error) {

	user, err := authService.AuthRepository.SignInWithGoogle(ctx, model)
	if err != nil {
		return entities.User{}, err
	}

	if err != nil {
		return entities.User{}, errors.New("invalid password")
	}

	return user, nil
}
