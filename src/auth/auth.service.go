package Auth

import (
	"FM/src/auth/models"
	"FM/src/entities"
	"context"
)

type AuthService interface {
	SignInWithGoogle(ctx context.Context, req models.SignInWithGoogleModles) (entities.User, error)
}
