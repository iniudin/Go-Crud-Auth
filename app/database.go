package app

import (
	"fmt"
	"go-crud-auth/config"
	"go-crud-auth/helper"
	"go-crud-auth/model/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase(config *config.Config) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Database.Host,
		config.Database.Port,
		config.Database.Username,
		config.Database.Password,
		config.Database.Name,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	helper.PanicError(err)

	db.AutoMigrate(
		&domain.User{},
		&domain.Mentor{},
		&domain.Course{},
		&domain.Video{},
		&domain.Attachment{},
	)

	return db
}
