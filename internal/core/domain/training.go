package domain

import (
	"time"

	"gorm.io/gorm"
)

type Training struct {
	gorm.Model

	Title string `json:"title" gorm:"not null;unique;size:255"`
	Description string `json:"description" gorm:"not null;size:255"`
	Date time.Time `json:"date" gorm:"not null;size:255"` // start date
	Duration int `json:"duration" gorm:"not null"` // duration in days
	Instructor string `json:"instructor" gorm:"not null;size:255"`

	CategoryID uint `json:"category_id"`
	Category Category `json:"category" gorm:"foreignKey:CategoryID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // belongs to category
}
