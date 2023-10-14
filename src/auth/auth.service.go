package Auth

import (
	"context"
	"FM/src/auth/models"
	"FM/src/entities"
)

type AuthService interface {

	SignInWithGoogle(ctx context.Context, model models.SignInWithGoogleModel) (entities.User, error)
}
