package app

import (
	"fmt"
	"go-crud-auth/config"
	"go-crud-auth/helper"

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
	return db
}
