package user

import (
	"FM/src/configuration"
	"FM/src/core/exception"
	"FM/src/core/http"
	"FM/src/core/middleware"
	"FM/src/core/shared"
	"FM/src/core/utils"
	modelUser "FM/src/modules/user/model"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type UserHandler struct {
	UserService
	configuration.Config
}

func NewCategoryHandler(userService *UserService, config configuration.Config) *UserHandler {
	return &UserHandler{UserService: *userService, Config: config}
}

func (handler UserHandler) Route(app *fiber.App) {
	var basePath = utils.GetBaseRoute(handler.Config, "/user")

	route := app.Group(basePath, middleware.AuthMiddleware(handler.Config), middleware.RoleMiddleware([]string{"TEACHER"}))
	route.Get("/", handler.GetProfile)
	route.Put("/", handler.UpdateProfile)
}

func (handler UserHandler) GetProfile(c *fiber.Ctx) error {
	userData := c.Locals("user")
	user := userData.(jwt.MapClaims)
	response, err := handler.UserService.GetProfile(c.Context(), uint(user["id"].(float64)))
	if err != nil {
		return exception.HandleError(c, err)
	}
	return c.Status(fiber.StatusOK).JSON(http.HttpResponse{
		StatusCode: fiber.StatusOK,
		Message:    "Get profile successfully",
		Data:       response,
	})
}

func (handler UserHandler) UpdateProfile(c *fiber.Ctx) error {
	 validator := shared.NewValidator()
	userData := c.Locals("user")
	user := userData.(jwt.MapClaims)
	var request modelUser.UpdateUserReq
	if err := c.BodyParser(&request); err != nil {
		return exception.HandleError(c, err)
	}

	if err := validator.Validate(request); err != nil {
		return exception.HandleErrorCustomMessage(c, "Missing required fields")
	}

	response, err := handler.UserService.UpdateProfile(c.Context(), uint(user["id"].(float64)), request)
	if err != nil {
		return exception.HandleError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(http.HttpResponse{
		StatusCode: fiber.StatusOK,
		Message:    "Update profile successfully",
		Data:       response,
	})
}
