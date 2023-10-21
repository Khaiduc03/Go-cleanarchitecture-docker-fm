package category

import (
	"FM/src/configuration"
	"FM/src/core/exception"
	"FM/src/core/http"
	"FM/src/core/utils"
	"FM/src/modules/category/model"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type CategoryHandler struct {
	CategoryService
	configuration.Config
}

func NewCategoryHandler(categoryService *CategoryService, config configuration.Config) *CategoryHandler {
	return &CategoryHandler{CategoryService: *categoryService, Config: config}
}

func (handler CategoryHandler) Route(app *fiber.App) {
	var basePath = utils.GetBaseRoute(handler.Config, "/category")

	route := app.Group(basePath)

	route.Get("/", handler.FindAll)
	route.Get("/:id", handler.FindById)
	route.Post("/", handler.Create)
	route.Put("/", handler.Update)
	route.Delete("/:id", handler.Delete)
}

func (handler CategoryHandler) FindAll(c *fiber.Ctx) error {
	categories, err := handler.CategoryService.FindAll(c.Context())
	if err != nil {
		return exception.HandleError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(http.HttpResponse{
		StatusCode: fiber.StatusOK,
		Message:    "Get all category successfully",
		Data:       categories,
	})
}

func (handler CategoryHandler) FindById(c *fiber.Ctx) error {
	idStr := c.Params("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return exception.HandleError(c, err)
	}

	category, err := handler.CategoryService.FindById(c.Context(), id)
	if err != nil {
		return exception.HandleError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(http.HttpResponse{
		StatusCode: fiber.StatusOK,
		Message:    "Get category by id successfully",
		Data:       category,
	})
}

func (handler CategoryHandler) Create(c *fiber.Ctx) error {
	var request model.CreateCategoryReq
	if err := c.BodyParser(&request); err != nil {
		return exception.HandleError(c, err)
	}

	message, err := handler.CategoryService.Create(c.Context(), request.Name)
	if err != nil {
		return exception.HandleError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(http.HttpResponse{
		StatusCode: fiber.StatusOK,
		Message:    message,
	})
}

func (handler CategoryHandler) Update(c *fiber.Ctx) error {
	var model model.UpdateCategoryReq
	if err := c.BodyParser(&model); err != nil {
		return exception.HandleError(c, err)
	}
	message, err := handler.CategoryService.Update(c.Context(), model)
	if err != nil {
		return exception.HandleError(c, err)
	}
	return c.Status(fiber.StatusOK).JSON(http.HttpResponse{
		StatusCode: fiber.StatusOK,
		Message:    message,
	})
}

func (handler CategoryHandler) Delete(c *fiber.Ctx) error {
	idStr := c.Params("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return exception.HandleError(c, err)
	}

	message, err := handler.CategoryService.Delete(c.Context(), id)
	if err != nil {
		return exception.HandleError(c, err)
	}
	return c.Status(fiber.StatusOK).JSON(http.HttpResponse{
		StatusCode: fiber.StatusOK,
		Message:    message,
	})
}
