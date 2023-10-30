package main

import (
	Auth "FM/src/auth"
	AuthImpl "FM/src/auth/implements"
	"FM/src/configuration"
	"FM/src/core/exception"
	"FM/src/core/firebase"

	room "FM/src/modules/Room"
	roomImpl "FM/src/modules/Room/implements"
	"FM/src/modules/category"
	categoryImpl "FM/src/modules/category/implements"
	"FM/src/modules/feedback"
	feedbackImpl "FM/src/modules/feedback/implements"
	"FM/src/modules/user"
	userImpl "FM/src/modules/user/implements"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	config := configuration.NewConfig()
	database := configuration.NewDataBase(config)
	// cloud, _err := libs.NewCloudinary(config)
	// if _err != nil {
	// 	panic(_err)
	// }
	firebaseApp := firebase.InitFirebaseAdmin()
	firebaseAuth := firebase.NewFirebaseAuth(&firebaseApp)

	//auth
	authRepository := AuthImpl.NewAuthRepositoryImpl(database)
	authService := AuthImpl.NewAuthServiceImpl(&authRepository, &firebaseAuth)
	authHandler := Auth.NewAuthHandler(&authService, config)

	//category
	categoryRepository := categoryImpl.NewCategoryRepositoryImpl(database)
	categoryService := categoryImpl.NewCategoryServiceImpl(&categoryRepository)
	categoryHandler := category.NewCategoryHandler(&categoryService, config)

	//user
	userRepository := userImpl.NewUserRepositoryImpl(database)
	userService := userImpl.NewUserServiceImpl(&userRepository)
	userHandler := user.NewCategoryHandler(&userService, config)

	//room
	roomRepository := roomImpl.NewRoomRepositoryImpl(database)
	roomService := roomImpl.NewRoomServiceImpl(&roomRepository)
	roomHandler := room.NewRoomHandler(&roomService, config)

	//feedback
	feedbackRepository := feedbackImpl.NewFeedbackRepositoryImpl(database)
	feedbackService := feedbackImpl.NewFeedbackServiceImpl(&feedbackRepository)
	feedbackHandler := feedback.NewFeedbackHandler(&feedbackService, config)

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

	// app.Use(cache.New(cache.Config{
	// 	Next:         nil,
	// 	Expiration:  1,
	// 	CacheHeader:  "X-Cache",
	// 	CacheControl: false,
	// 	KeyGenerator: func(c *fiber.Ctx) string {
	// 		return utils.CopyString(c.Path())
	// 	},
	// 	ExpirationGenerator:  nil,
	// 	StoreResponseHeaders: false,
	// 	Storage:              nil,
	// 	MaxBytes:             0,
	// 	Methods:              []string{fiber.MethodGet, fiber.MethodHead},
	// }))
	authHandler.Route(app)
	categoryHandler.Route(app)
	userHandler.Route(app)
	roomHandler.Route(app)
	feedbackHandler.Route(app)
	err := app.Listen(config.Get("SERVER_PORT"))

	exception.PanicLogging(err)
}
