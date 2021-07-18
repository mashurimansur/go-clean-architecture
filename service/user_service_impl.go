package service

import (
	"clean-arch-go/entity"
	"clean-arch-go/model"
	"clean-arch-go/repository"

	validation "github.com/go-ozzo/ozzo-validation"
)

type userServiceImpl struct {
	UserRepository repository.UserRepository
}

func NewUserService(userRepository *repository.UserRepository) UserService {
	return &userServiceImpl{UserRepository: *userRepository}
}

func (service *userServiceImpl) List() (response []model.UserResponse, err error) {
	users, err := service.UserRepository.List()

	for _, v := range users {
		res := model.UserResponse{
			ID:        v.ID,
			Name:      v.Name,
			Email:     v.Email,
			Age:       v.Age,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}

		response = append(response, res)
	}

	return

}

func (service *userServiceImpl) Create(request model.UserRequest) error {
	errValidation := validation.Validate(request)
	if errValidation != nil {
		return errValidation
	}

	user := entity.User{
		Name:  request.Name,
		Email: request.Email,
		Age:   request.Age,
	}

	errCreate := service.UserRepository.Create(user)

	return errCreate
}
