package domain

import "gorm.io/gorm"

type Video struct {
	gorm.Model
	Title        string
	Url string
	Description string
	CourseID    uint
}