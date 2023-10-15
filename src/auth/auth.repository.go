package Auth

import (
	"context"
	"FM/src/auth/models"
	"FM/src/entities"
)

type AuthRepository interface {

	SignInWithGoogle(ctx context.Context, model models.Payload) (entities.User, error)
}
