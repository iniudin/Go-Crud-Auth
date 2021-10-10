package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name        string
	Username string
	Email string
	Password string
	Role    uint
}