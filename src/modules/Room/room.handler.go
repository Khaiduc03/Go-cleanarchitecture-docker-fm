package room

import (
	"FM/src/configuration"
	"FM/src/core/exception"
	"FM/src/core/http"
	"FM/src/core/middleware"
	"FM/src/core/shared"
	"FM/src/core/utils"
	modelRoom "FM/src/modules/Room/model"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type RoomHandler struct {
	RoomService RoomService
	configuration.Config
}

func NewRoomHandler(roomService *RoomService, config configuration.Config) *RoomHandler {
	return &RoomHandler{RoomService: *roomService, Config: config}
}

func (handler RoomHandler) Route(app *fiber.App) {
	var basePath = utils.GetBaseRoute(handler.Config, "/room")

	route := app.Group(basePath, middleware.AuthMiddleware(handler.Config), middleware.RoleMiddleware([]string{"TEACHER"}))
	route.Get("/", handler.FindAll)
	route.Get("/:id", handler.FindById)
	route.Post("/", handler.Create)
	route.Put("/", handler.Update)
	route.Delete("/:id", handler.Delete)
}

func (handler RoomHandler) FindAll(c *fiber.Ctx) error {
	rooms, err := handler.RoomService.FindAll(c.Context())
	if err != nil {
		return exception.HandleError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(http.HttpResponse{
		StatusCode: fiber.StatusOK,
		Message:    "Get all room successfully",
		Data:       rooms,
	})
}

func (handler RoomHandler) FindById(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return exception.HandleError(c, err)
	}

	room, err := handler.RoomService.FindById(c.Context(), id)
	if err != nil {
		return exception.HandleError(c, err)
	}
	if room.ID <= 0 {
		return exception.HandleErrorCustomMessage(c, "Room not found")
	}

	return c.Status(fiber.StatusOK).JSON(http.HttpResponse{
		StatusCode: fiber.StatusOK,
		Message:    "Get room by id successfully",
		Data:       room,
	})
}

func (handler RoomHandler) Create(c *fiber.Ctx) error {
	 validator := shared.NewValidator()
	var req modelRoom.CreateRoomReq
	if err := c.BodyParser(&req); err != nil {
		return exception.HandleError(c, err)
	}


	if err := validator.Validate(req); err != nil {
		return exception.HandleErrorCustomMessage(c, "Missing required fields")
	}
	// return c.Status(fiber.StatusOK).JSON(http.HttpResponse{
	// 	StatusCode: fiber.StatusOK,
	// 	Message:    "ok",
	// 	Data: req,
	// })
	message, err := handler.RoomService.Create(c.Context(), req)
	if err != nil {
		return exception.HandleError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(http.HttpResponse{
		StatusCode: fiber.StatusOK,
		Message:    message,
	})
}

func (handler RoomHandler) Update(c *fiber.Ctx) error {
	validator := shared.NewValidator()
	var req modelRoom.UpdateRoomReq
	if err := c.BodyParser(&req); err != nil {
		return exception.HandleError(c, err)
	}

	if err := validator.Validate(req); err != nil {
		return exception.HandleErrorCustomMessage(c, err.Error())
	}

	message, err := handler.RoomService.Update(c.Context(), req)
	if err != nil {
		return exception.HandleError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(http.HttpResponse{
		StatusCode: fiber.StatusOK,
		Message:    message,
		Data: "ok",
	})
}

func (handler RoomHandler) Delete(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return exception.HandleError(c, err)
	}

	message, err := handler.RoomService.Delete(c.Context(), id)
	if err != nil {
		return exception.HandleError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(http.HttpResponse{
		StatusCode: fiber.StatusOK,
		Message:    message,
	})
}
