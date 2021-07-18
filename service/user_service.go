package service

import "clean-arch-go/entity"

type UserService interface {
	List() (response []entity.User, err error)
	Create(request entity.User) (err error)
}
