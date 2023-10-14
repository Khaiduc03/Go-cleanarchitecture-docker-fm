package firebase

import (
	"FM/src/core/exception"
	"context"
	"path"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
)

func InitFirebaseAdmin() firebase.App {
	path := path.Join("firebse.json")
	ctx := context.Background()
	opt := option.WithCredentialsFile(path)
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		exception.PanicLogging(err.Error())
	}
	return *app
}

type FirebaseAuth struct {
	auth.Client
}

func NewFirebaseAuth(firebase *firebase.App) FirebaseAuth {
	client, err := firebase.Auth(context.Background())
	if err != nil {
		exception.PanicLogging(err)
	}
	return FirebaseAuth{Client: *client}
}

type UserPayload struct {
	Email   string `json:"email"`
	UserID  string `json:"user_id"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
}
