package repository

import (
	"clean-arch-go/entity"

	"gorm.io/gorm"
)

type userRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(database *gorm.DB) UserRepository {
	return &userRepositoryImpl{db: database}
}

func (repository *userRepositoryImpl) List() (users []entity.User, err error) {
	if err := repository.db.Find(&users).Error; err != nil {
		return nil, err
	}

	return
}

func (repository *userRepositoryImpl) Create(request entity.User) (err error) {
	if err := repository.db.Model(&request).Create(&request).Error; err != nil {
		return err
	}

	return
}
