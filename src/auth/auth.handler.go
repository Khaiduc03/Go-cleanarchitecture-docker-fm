package Auth

import (
	"FM/src/auth/models"
	"FM/src/configuration"
	"FM/src/core/exception"
	"FM/src/core/http"
	"FM/src/core/utils"
	"fmt"

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

	route.Post("/login", handler.SignInWithGoogle)
}

func (handler AuthHandler) SignInWithGoogle(c *fiber.Ctx) error {

	if c.Method() != "POST" {
		return c.Status(fiber.StatusMethodNotAllowed).JSON(http.HttpResponse{
			StatusCode: fiber.StatusMethodNotAllowed,
			Message:    "Method Not Allowed",
		})
	}

	var requestData = models.SignInWithGoogleModles{}

	if err := c.BodyParser(&requestData); err != nil {
		exception.HandleError(c, err)
	}

	result, err := handler.AuthService.SignInWithGoogle(c.Context(), requestData)
	//fmt.Println(result)
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(http.HttpResponse{
			StatusCode: fiber.ErrBadRequest.Code,
			Message:    "Sign in with google failed",
			Data:       err,
		})
	}

	//check result is empty
	fmt.Println(result)

	return c.Status(fiber.StatusOK).JSON(http.HttpResponseNoData{
		StatusCode: fiber.StatusOK,
		Message:    "Sign in with google successfully",
	})
}
