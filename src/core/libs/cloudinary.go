package libs

import (
	"bytes"
	"context"
	"log"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func UploadCloudinary(ctx context.Context,file []byte) string {
	cld, err := cloudinary.NewFromURL("cloudinary://648127967397929:Wgwk-Gb5c8nFLVsMijmPnPiQnNY@dzycibpuc")
	if err != nil {
		log.Fatalf("Failed to create cloudinary instance, %v\n", err)
	}

	uploadResult, err := cld.Upload.Upload(
		ctx,
		bytes.NewReader(file),
		uploader.UploadParams{PublicID: "logo"})
	if err != nil {
		log.Fatalf("Failed to upload file, %v\n", err)
	}
	log.Println(uploadResult.SecureURL)
	return uploadResult.URL
}
