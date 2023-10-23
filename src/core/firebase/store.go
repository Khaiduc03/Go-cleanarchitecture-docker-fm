package firebase

import (
	"FM/src/core/exception"
	"context"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/storage"
)

type FirebaseStore struct {
	storage.Client
}

func NewFirebaseStore(firebase *firebase.App) FirebaseStore {
	client, err := firebase.Storage(context.Background())
	if err != nil {
		exception.PanicLogging(err)
	}
	return FirebaseStore{Client: *client}
}

func (store FirebaseStore) UploadFile(ctx context.Context, file []byte, fileName string) (string, error) {
	bucket, err := store.Client.DefaultBucket()
	if err != nil {
		return "", err
	}

	obj := bucket.Object(fileName)
	w := obj.NewWriter(ctx)
	if _, err := w.Write(file); err != nil {
		return "", err
	}
	if err := w.Close(); err != nil {
		return "", err
	}
	attrs, err := obj.Attrs(ctx)
	if err != nil {
		return "", err
	}

	url := attrs.MediaLink
	return url, nil
}
