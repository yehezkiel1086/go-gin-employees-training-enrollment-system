package domain

import (
	"time"

	"gorm.io/gorm"
)

type Enrollment struct {
	gorm.Model

	UserID     uint      `json:"user_id"`
	User       User      `json:"user" gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // belongs to user

	TrainingID uint      `json:"training_id"`
	Training   Training  `json:"training" gorm:"foreignKey:TrainingID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // belongs to training

	EnrolledAt time.Time `json:"enrolled_at"`
}
