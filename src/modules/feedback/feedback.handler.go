package feedback

import (
	"FM/src/configuration"
	"FM/src/core/exception"
	"FM/src/core/http"
	"FM/src/core/libs"
	"FM/src/core/middleware"
	"FM/src/core/utils"
	modelFeedback "FM/src/modules/feedback/model"
	"context"
	"io"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type FeedbackHandler struct {
	FeedbackService FeedBackService
	configuration.Config
}

func NewFeedbackHandler(feedbackService *FeedBackService, config configuration.Config) *FeedbackHandler {

	return &FeedbackHandler{FeedbackService: *feedbackService, Config: config}
}

func (handler FeedbackHandler) Route(app *fiber.App) {
	var basePath = utils.GetBaseRoute(handler.Config, "/feedback")

	route := app.Group(basePath, middleware.AuthMiddleware(handler.Config), middleware.RoleMiddleware([]string{"TEACHER"}))
	route.Get("/", handler.FindAll)
	route.Get("/:id", handler.FindById)
	route.Post("/", handler.Create)
	// route.Put("/", handler.Update)
	// route.Delete("/:id", handler.Delete)
}

func (handler FeedbackHandler) FindAll(c *fiber.Ctx) error {
	feedbacks, err := handler.FeedbackService.FindAll(c.Context())
	if err != nil {
		return exception.HandleError(c, err)
	}
	return c.Status(fiber.StatusOK).JSON(http.HttpResponse{
		StatusCode: fiber.StatusOK,
		Message:    "Get all feed successfully",
		Data:       feedbacks,
	})
}

// func (handler FeedBackService) HistoryFeedback(c *fiber.Ctx) error{

// }

// find by id
func (handler FeedbackHandler) FindById(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return exception.HandleError(c, err)
	}

	feedback, err := handler.FeedbackService.FindById(c.Context(), id)
	if err != nil {
		return exception.HandleError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(http.HttpResponse{
		StatusCode: fiber.StatusOK,
		Message:    "Get feedback by id successfully",
		Data:       feedback,
	})
}

// create feedback
func (handler FeedbackHandler) Create(c *fiber.Ctx) error {
	//data of user in request
	userData := c.Locals("user")
	user := userData.(jwt.MapClaims)

	var urls []string
	var req modelFeedback.CreateFeedbackReq
	form, err := c.MultipartForm()
	if err != nil {
		return exception.HandleError(c, err)
	}

	categoryID, err := strconv.ParseUint(form.Value["category_id"][0], 10, 32)
	if err != nil {
		return exception.HandleError(c, err)
	}
	checkCategory := handler.FeedbackService.CheckCategory(c.Context(), int(categoryID))
	if checkCategory != nil {
		return exception.HandleErrorCustomMessage(c, "Category not found")
	}
	roomID, err := strconv.ParseUint(form.Value["room_id"][0], 10, 32)
	if err != nil {
		return exception.HandleError(c, err)
	}
	checkRoom := handler.FeedbackService.CheckRoom(c.Context(), int(roomID))
	if checkRoom != nil {
		return exception.HandleErrorCustomMessage(c, "Room not found")
	}

	files := form.File["files"]
	if len(files) == 0 {
		return exception.HandleErrorCustomMessage(c, "Missing required fields")
	}
	//upload to cloud and get url
	for _, file := range files {
		res, err := file.Open()
		if err != nil {
			return exception.HandleError(c, err)
		}
		fileByte, err := io.ReadAll(res)
		if err != nil {
			return exception.HandleError(c, err)
		}
		url := libs.UploadCloudinary(context.Background(), fileByte)
		urls = append(urls, url)
	}

	req = modelFeedback.CreateFeedbackReq{
		UserID:         uint(user["id"].(float64)),
		Name_Feed_Back: form.Value["name_feed_back"][0],
		Description:    form.Value["description"][0],
		RoomID:         uint(roomID),
		CategoryID:     uint(categoryID),
		Urls:           urls,
	}
	//fmt.Println("req", req)

	if err := c.BodyParser(&req); err != nil {
		return exception.HandleError(c, err)
	}

	message, err := handler.FeedbackService.Create(c.Context(), req)
	if err != nil {
		return exception.HandleError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(http.HttpResponse{
		StatusCode: fiber.StatusOK,
		Message:    message,
	})
}
