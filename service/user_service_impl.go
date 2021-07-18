package service

import (
	"clean-arch-go/entity"
	"clean-arch-go/repository"
)

type userServiceImpl struct {
	UserRepository repository.UserRepository
}

func NewUserService(userRepository *repository.UserRepository) UserService {
	return &userServiceImpl{UserRepository: *userRepository}
}

func (service *userServiceImpl) List() (response []entity.User, err error) {
	response, err = service.List()

	return

}

func (service *userServiceImpl) Create(request entity.User) (err error) {
	// validation.Validate(request)

	err = service.UserRepository.Create(request)

	return
}
