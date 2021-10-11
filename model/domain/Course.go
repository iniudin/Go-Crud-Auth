package domain

import "gorm.io/gorm"

type Course struct {
	gorm.Model
	Title       string
	Description string
	Thumbnail   string
	MentorID    uint
}
