package libs

import (
	"bytes"
	"context"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/google/uuid"
)

func UploadCloudinary(ctx context.Context, file []byte, filename string) string {
	uuid := uuid.New().String()
	cloudName := "dzycibpuc"
	apiKey := "648127967397929"
	apiSecret := "Wgwk-Gb5c8nFLVsMijmPnPiQnNY"

	cld, _ := cloudinary.NewFromParams(cloudName, apiKey, apiSecret)

	uploadResult, _ := cld.Upload.Upload(ctx, bytes.NewReader(file), uploader.UploadParams{PublicID: uuid, Folder: "images"})

	return uploadResult.URL

}
