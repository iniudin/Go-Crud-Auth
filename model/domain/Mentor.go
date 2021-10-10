package domain

import "gorm.io/gorm"

type Mentor struct {
	gorm.Model
	Occupation        string
	Institution string
	Photo string
	UserID    uint
}