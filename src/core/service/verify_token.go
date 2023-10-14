package firebase

import (
	"context"
)

func (auth FirebaseAuth) VerifyIDToken(ctx context.Context, idToken string) (UserPayload, error) {
	token, err := auth.Client.VerifyIDToken(ctx, idToken)
	if err != nil {
		return UserPayload{}, err
	}

	claims := token.Claims

	userID := claims["user_id"].(string)
	name := claims["name"].(string)
	email := claims["email"].(string)
	picture := claims["picture"].(string)

	return UserPayload{
		UserID:  userID,
		Name:    name,
		Email:   email,
		Picture: picture,
	}, nil
}
