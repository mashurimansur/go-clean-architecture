package service

import (
	"clean-arch-go/entity"
	"clean-arch-go/model"
)

type UserService interface {
	List() (response []entity.User, err error)
	Create(request model.UserRequest) error
}
