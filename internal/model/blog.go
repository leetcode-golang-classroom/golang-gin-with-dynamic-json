package model

import "gorm.io/gorm"

type Blog struct {
	gorm.Model
	Title   string `gorm:"size:255"`
	Content string `gorm:"type:text"`
}
