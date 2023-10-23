package firebase

import (
	"FM/src/core/exception"
	"context"
	"path"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

func InitFirebaseAdmin() firebase.App {
	path := path.Join("firebase.json")
	opt := option.WithCredentialsFile(path)
	config := &firebase.Config{
		StorageBucket: "fmanager-795a5.appspot.com",
	}
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		exception.PanicLogging(err.Error())
	}
	return *app
}
