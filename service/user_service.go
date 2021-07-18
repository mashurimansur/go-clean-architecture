package service

import (
	"clean-arch-go/model"
)

type UserService interface {
	List() (response []model.UserResponse, err error)
	Create(request model.UserRequest) error
}
