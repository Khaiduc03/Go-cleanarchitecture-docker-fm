package libs

import (
	"FM/src/configuration"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func GoogleConfig(config configuration.Config) *oauth2.Config {
	CLIENT_ID := config.Get("CLIENT_ID")
	googleOAuthConfig := &oauth2.Config{
		ClientID: CLIENT_ID,
		Scopes:   []string{"https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint: google.Endpoint,
	}

	return googleOAuthConfig
}
