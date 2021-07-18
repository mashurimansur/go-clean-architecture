package controller

import (
	"clean-arch-go/model"
	"clean-arch-go/service"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	UserService service.UserService
}

func NewUserController(userService *service.UserService) UserController {
	return UserController{UserService: *userService}
}

func (controller *UserController) Route(app *fiber.App) {
	app.Get("/api/users", controller.List)
	app.Post("/api/users", controller.Create)
}

func (controller *UserController) List(c *fiber.Ctx) error {
	users, err := controller.UserService.List()
	if err != nil {
		return c.JSON(model.WebResponse{
			Code:   404,
			Status: err.Error(),
		})
	}

	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   users,
	})
}

func (controller *UserController) Create(c *fiber.Ctx) error {
	var request model.UserRequest
	err := c.BodyParser(&request)
	if err != nil {
		return c.JSON(model.WebResponse{
			Code:   400,
			Status: err.Error(),
		})
	}

	errCreate := controller.UserService.Create(request)
	if errCreate != nil {
		return c.JSON(model.WebResponse{
			Code:   400,
			Status: err.Error(),
		})
	}

	return c.JSON(model.WebResponse{
		Code:   201,
		Status: "Create User Success",
	})
}
