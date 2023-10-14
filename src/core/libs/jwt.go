package libs

import (
	"FM/src/configuration"
	"FM/src/core/exception"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const (
	AccessToken  = 0
	RefreshToken = 1
) // token type

type JWTPayload struct {
	ID          uint   `json:"id"`
	Email       string `json:"email"`
	Role        string `json:"role"`
	DeviceToken string `json:"device_token"`
}

func GenerateToken(payload JWTPayload, tokenType int, config configuration.Config) string {
	accessTokenSecret := config.Get("ACCESS_TOKEN_SECRET")
	refreshTokenSecret := config.Get("REFRESH_TOKEN_SECRET")

	accessTokenExpired, err := strconv.Atoi(config.Get("ACCESS_TOKEN_EXPIRE_MINUTES_COUNT"))
	exception.PanicLogging(err)
	refreshTokenExpired, err := strconv.Atoi(config.Get("REFRESH_TOKEN_EXPIRE_MINUTES_COUNT"))
	exception.PanicLogging(err)

	claims := jwt.MapClaims{}

	claims["id"] = payload.ID
	claims["email"] = payload.Email
	claims["role"] = payload.Role

	var jwtSecret string
	var expired int64

	if tokenType == AccessToken {
		jwtSecret = accessTokenSecret
		expired = time.Now().Add(time.Minute * time.Duration(accessTokenExpired)).Unix()
	} else if tokenType == RefreshToken {
		jwtSecret = refreshTokenSecret
		expired = time.Now().Add(time.Minute * time.Duration(refreshTokenExpired)).Unix()
	}

	claims["exp"] = expired

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenSigned, err := token.SignedString([]byte(jwtSecret))
	exception.PanicLogging(err)

	return tokenSigned
}

func VerifyToken(tokenString string, tokenType int, config configuration.Config) (jwt.MapClaims, error) {
	accessTokenSecret := config.Get("ACCESS_TOKEN_SECRET")
	refreshTokenSecret := config.Get("REFRESH_TOKEN_SECRET")

	var jwtSecret string

	if tokenType == AccessToken {
		jwtSecret = accessTokenSecret
	} else if tokenType == RefreshToken {
		jwtSecret = refreshTokenSecret
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid {
		return nil, err
	}

	return claims, nil
}
