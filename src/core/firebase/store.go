package firebase

import (
	"FM/src/core/exception"
	"context"
	"log"

	Cloudstorage "cloud.google.com/go/storage"
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
		log.Fatalf("Error getting default bucket: %v\n", err)
		return "", err
	}
	wc := bucket.Object(fileName).NewWriter(ctx)
	wc.ContentType = "image/jpeg"

	wc.ACL = []Cloudstorage.ACLRule{{Entity: Cloudstorage.AllUsers, Role: Cloudstorage.RoleReader}}

	if _, err := wc.Write([]byte(file)); err != nil {
		log.Fatalf("Error writing to storage: %v\n", err)
		return "", err
	}
	if err := wc.Close(); err != nil {
		log.Fatalf("Error writing to storage: %v\n", err)
		return "", err
	}
	return wc.Attrs().MediaLink, nil
}
