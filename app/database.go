package app

import (
	"Secreto/helper"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase() *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		"localhost", "5432", "postgres", "Klopbgt12!", "secreto_db")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	helper.PanicError(err)
	return db
}
