package Auth

import (
	"FM/src/auth/models"
	"FM/src/configuration"
	"FM/src/core/exception"
	"FM/src/core/http"
	"FM/src/core/libs"
	firebase "FM/src/core/service"
	"FM/src/core/utils"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	AuthService
	configuration.Config
	firebase.FirebaseAuth
}

func NewAuthHandler(authService *AuthService, config configuration.Config, firebaseAuth *firebase.FirebaseAuth) *AuthHandler {
	return &AuthHandler{AuthService: *authService, Config: config, FirebaseAuth: *firebaseAuth}
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

	var requestData struct {
		IDToken string `json:"idToken"`
	}

	if err := c.BodyParser(&requestData); err != nil {
		exception.HandleError(c, err)
	}

	idToken := requestData.IDToken

	claims, err := handler.FirebaseAuth.VerifyIDToken(c.Context(), idToken)
	exception.HandleError(c, err)

	model := models.SignInWithGoogleModel{
		UserID: claims.UserID,
		Email:  claims.Email,
	}

	result, err := handler.AuthService.SignInWithGoogle(c.Context(), model)
	exception.HandleError(c, err)

	payload := libs.JWTPayload{
		ID:          result.ID,
		Email:       result.Email,
		Role:        result.Role,
		DeviceToken: result.DeviceToken,
	}

	accessToken := libs.GenerateToken(payload, libs.AccessToken, handler.Config)
	refreshToken := libs.GenerateToken(payload, libs.RefreshToken, handler.Config)

	response := models.ResponseSignIn{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return c.Status(fiber.StatusOK).JSON(http.HttpResponse{
		StatusCode: fiber.StatusOK,
		Message:    "Sign in with google successfully",
		Data:       response,
	})
}
