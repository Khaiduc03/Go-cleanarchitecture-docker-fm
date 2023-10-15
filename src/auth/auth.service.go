package Auth

import (
	"FM/src/auth/models"
	"context"
)

type AuthService interface {
	SignInWithGoogle(ctx context.Context, idToken string) (models.Payload, error)
}
