package exception

import (
	"FM/src/core/http"

	"github.com/gofiber/fiber/v2"
)

func PanicLogging(err interface{}) {
	if err != nil {
		panic(err)
	}
}

func HandleError(c *fiber.Ctx, err error) error {
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(http.HttpResponse{
			StatusCode: fiber.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	return nil
}
