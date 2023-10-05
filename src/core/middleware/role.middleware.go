package middleware

import (
	"FM/src/core/http"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func RoleMiddleware(role []string) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		user := c.Locals("user").(jwt.MapClaims)

		userRole := user["role"].(string)

		for _, r := range role {
			if r == userRole {
				return c.Next()
			}
		}

		return c.Status(fiber.StatusOK).JSON(http.HttpResponse{
			StatusCode: fiber.StatusForbidden,
			Message:    "Unauthorized",
			Data:       nil,
		})
	}
}
