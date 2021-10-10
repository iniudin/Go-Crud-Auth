package domain

import "gorm.io/gorm"

type Attachment struct {
	gorm.Model
	Title        string
	Attachment string
	Type    uint
	VideoID    uint
}