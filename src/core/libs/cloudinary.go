package libs

import (
	"FM/src/configuration"
	"FM/src/core/exception"
	"context"

	"github.com/cloudinary/cloudinary-go/v2"
)

func NewCloudinary(config configuration.Config) (*cloudinary.Cloudinary, context.Context) {

	CLOUD_NAME := config.Get("CLOUD_NAME")
	API_KEY := config.Get("API_KEY")
	API_SECRET := config.Get("API_SECRET")

	cld, err := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)
	exception.PanicLogging(err)

	cld.Config.URL.Secure = true
	var ctx = context.Background()

	return cld, ctx
}
