package repository

import "clean-arch-go/entity"

type UserRepository interface {
	List() (users []entity.User, err error)
	Create(request entity.User) (err error)
}
