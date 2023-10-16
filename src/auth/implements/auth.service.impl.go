package AuthImpl

import (
	Auth "FM/src/auth"
	"FM/src/auth/models"
	firebase "FM/src/core/service"
	"context"
	"fmt"
)

type AuthServiceImpl struct {
	Auth.AuthRepository
	firebase.FirebaseAuth
}

func NewAuthServiceImpl(authRepository *Auth.AuthRepository, firebaseAuth *firebase.FirebaseAuth) Auth.AuthService {
	return &AuthServiceImpl{AuthRepository: *authRepository, FirebaseAuth: *firebaseAuth}
}

func (authService *AuthServiceImpl) SignInWithGoogle(ctx context.Context, idToken string) (models.Payload, error) {

	claims, err := authService.FirebaseAuth.VerifyIDToken(ctx, idToken)
	if err != nil {
		fmt.Println(err)
	}

	payload := models.Payload{
		Name:    claims.Name,
		Email:   claims.Email,
		UserID:  claims.UserID,
		Picture: claims.Picture,
	}
	return payload, err

}
