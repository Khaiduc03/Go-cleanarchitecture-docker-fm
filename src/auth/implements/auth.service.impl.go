package AuthImpl

import (
	Auth "FM/src/auth"
	models "FM/src/auth/models"
	"FM/src/core/firebase"
	"FM/src/entities"
	"context"
	"fmt"

	"github.com/golang-jwt/jwt/v4"
)

type AuthServiceImpl struct {
	Auth.AuthRepository
	firebase.FirebaseAuth
}

func NewAuthServiceImpl(authRepository *Auth.AuthRepository, firebaseAuth *firebase.FirebaseAuth) Auth.AuthService {
	return &AuthServiceImpl{AuthRepository: *authRepository, FirebaseAuth: *firebaseAuth}
}

func (authService *AuthServiceImpl) SignInWithGoogle(ctx context.Context, req models.SignInWithGoogleModles) (entities.User, error) {
	idToken := req.IDToken
	position := req.Position

	if position == "HCM" || position == "HN" {
		position = "FPT Polytechnic Hồ Chí Minh"
	}

	tokenString, err := jwt.Parse(idToken, func(token *jwt.Token) (interface{}, error) {
		return []byte("18723735524-8qe3014rf4goh1ck0o6lp07tn7c0965q.apps.googleusercontent.com"), nil
	})
	fmt.Print(err)

	if err.Error() != "key is of invalid type" {
		return entities.User{}, err
	}

	claims := tokenString.Claims.(jwt.MapClaims)

	payload := models.Payload{
		Name:     claims["name"].(string),
		Email:    claims["email"].(string),
		Picture:  claims["picture"].(string),
		Position: position,
	}

	return authService.AuthRepository.SignInWithGoogle(ctx, payload)
}
