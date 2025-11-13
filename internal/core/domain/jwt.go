package domain

import "github.com/golang-jwt/jwt/v4"

type JWT struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Role Role `json:"role"`

	jwt.RegisteredClaims
}
