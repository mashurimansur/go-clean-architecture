package main

import (
	"clean-arch-go/config"
	"clean-arch-go/controller"
	"clean-arch-go/exception"
	"clean-arch-go/repository"
	"clean-arch-go/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	//setup configuration
	configuration := config.New()
	database := config.NewPostgreConnection(configuration)

	//setup repository
	userRepository := repository.NewUserRepository(database)

	//setup service
	userService := service.NewUserService(&userRepository)

	// setup controller
	userController := controller.NewUserController(&userService)

	// setup fiber
	app := fiber.New(config.NewFiberConfig())
	app.Use(recover.New())

	//setup routing
	userController.Route(app)

	//start app
	err := app.Listen(":8000")
	exception.PanicIfNeeded(err)

}
