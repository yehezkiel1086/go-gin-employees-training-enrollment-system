package domain

import "gorm.io/gorm"

type Role uint16

const (
	ADMIN_ROLE Role = 5150
	USER_ROLE Role = 2001
)

type User struct {
	gorm.Model

	Name string `json:"name" gorm:"not null;size:255"`
	Email string `json:"email" gorm:"not null;size:255;unique"`
	Password string `json:"password" gorm:"not null;size:255"`
	Role Role `json:"role" gorm:"not null;default:2001"`
}
