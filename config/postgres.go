package config

import (
	"clean-arch-go/exception"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgreConnection(config Config) *gorm.DB {
	username := config.Get("PG_USERNAME")
	password := config.Get("PG_PASSWORD")
	database := config.Get("PG_DATABASE")
	host := config.Get("PG_HOST")
	port := config.Get("PG_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", host, username, password, database, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	exception.PanicIfNeeded(err)

	return db
}
