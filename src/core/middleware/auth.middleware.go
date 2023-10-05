package middleware

import (
	"FM/src/configuration"
	"FM/src/core/http"
	"FM/src/core/libs"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(config configuration.Config) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusOK).JSON(http.HttpResponse{
				StatusCode: fiber.StatusUnauthorized,
				Message:    "Unauthorized",
				Data:       nil,
			})
		}

		token := strings.Split(authHeader, " ")[1]

		if token == "" {
			return c.Status(fiber.StatusOK).JSON(http.HttpResponse{
				StatusCode: fiber.StatusUnauthorized,
				Message:    "Unauthorized",
				Data:       nil,
			})
		}

		payload, err := libs.VerifyToken(token, libs.AccessToken, config)
		if err != nil {
			return c.Status(fiber.StatusOK).JSON(http.HttpResponse{
				StatusCode: fiber.StatusUnauthorized,
				Message:    err.Error(),
				Data:       nil,
			})
		}

		c.Locals("user", payload)
		return c.Next()
	}
}
