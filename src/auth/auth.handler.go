package Auth

import (
	"FM/src/configuration"
	"FM/src/core/exception"
	"FM/src/core/http"
	"FM/src/core/utils"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	AuthService
	configuration.Config
}

func NewAuthHandler(authService *AuthService, config configuration.Config) *AuthHandler {
	return &AuthHandler{AuthService: *authService, Config: config}
}

func (handler AuthHandler) Route(app *fiber.App) {
	var basePath = utils.GetBaseRoute(handler.Config, "/auth")
	route := app.Group(basePath)

	route.Get("/login", handler.SignInWithGoogle)
}

func (handler AuthHandler) SignInWithGoogle(c *fiber.Ctx) error {

	idToken := c.Query("idToken")

	result, err := handler.AuthService.SignInWithGoogle(c.Context(), idToken)

	if err != nil {
		exception.HandleError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(http.HttpResponse{
		StatusCode: fiber.StatusOK,
		Message:    "Sign in with google successfully",
		Data:       result,
	})
}
