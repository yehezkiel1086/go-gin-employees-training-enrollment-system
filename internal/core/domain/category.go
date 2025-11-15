package domain

import "gorm.io/gorm"

type Category struct {
	gorm.Model

	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"unique;not null;size:255"`
}
