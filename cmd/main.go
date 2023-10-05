package main

import (
	"FM/src/configuration"
	"FM/src/core/exception"
	"FM/src/core/libs"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	config := configuration.NewConfig()
	 database := configuration.NewDataBase(config)
	cloud, _err := libs.NewCloudinary(config)
	if _err != nil {
		panic(_err)
	}

	app := fiber.New(configuration.NewFiberConfiguration())

	app.Use(recover.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: true,
		AllowMethods:     "GET, POST, PUT, DELETE, PATCH",
		MaxAge:           3600,
		Next:             nil,
	}))

	app.Use(logger.New(logger.Config{
		Format: "${ip}:${port} ${status} - ${method} ${path}\n",
	}))
	app.Use(limiter.New(limiter.Config{
		Max:               20,
		Expiration:        30 * time.Second,
		LimiterMiddleware: limiter.SlidingWindow{},
	}))
	app.Use(cache.New(cache.Config{
		Next: func(c *fiber.Ctx) bool {
			return strings.Contains(c.Route().Path, "/ws")
		},
	}))

	err := app.Listen(config.Get("SERVER_PORT"))

	exception.PanicLogging(err)
}
