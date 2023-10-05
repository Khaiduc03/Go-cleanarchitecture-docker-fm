package utils

import (
	"bytes"
	"FM/src/configuration"
)

func GetBaseRoute(config configuration.Config, prefix string) string {
	API_VERSION := config.Get("API_VERSION")
	var basePath bytes.Buffer
	basePath.WriteString("/api/")
	basePath.WriteString(API_VERSION)
	basePath.WriteString(prefix)

	return basePath.String()
}
