package controller

import (
	"bytes"
	"clean-arch-go/config"
	"clean-arch-go/model"
	"clean-arch-go/repository"
	"clean-arch-go/service"
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/stretchr/testify/assert"
)

func createTestApp() *fiber.App {
	var app = fiber.New(config.NewFiberConfig())
	app.Use(recover.New())

	userController.Route(app)

	return app
}

var app = createTestApp()

var database = config.NewSQLiteConnection()
var userRepository = repository.NewUserRepository(database)
var userService = service.NewUserService(&userRepository)
var userController = NewUserController(&userService)

func TestUserControllerList(t *testing.T) {
	request := httptest.NewRequest("GET", "/api/users", nil)
	request.Header.Set("Accept", "application/json")

	response, _ := app.Test(request)

	assert.Equal(t, 200, response.StatusCode)
	responseBody, _ := ioutil.ReadAll(response.Body)

	webResponse := model.WebResponse{}
	json.Unmarshal(responseBody, &webResponse)
	assert.Equal(t, 200, webResponse.Code)
	assert.Equal(t, "OK", webResponse.Status)

}

func TestUserControllerCreate(t *testing.T) {
	userRequest := model.UserRequest{
		Name:  "Test",
		Email: "test@mail.com",
		Age:   20,
	}

	requestBody, _ := json.Marshal(userRequest)
	request := httptest.NewRequest("POST", "/api/users", bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, _ := app.Test(request)
	assert.Equal(t, 200, response.StatusCode)
	responseBody, _ := ioutil.ReadAll(response.Body)

	webResponse := model.WebResponse{}
	json.Unmarshal(responseBody, &webResponse)
	assert.Equal(t, 201, webResponse.Code)
	assert.Equal(t, "Create User Success", webResponse.Status)
}
