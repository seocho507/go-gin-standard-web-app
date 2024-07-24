package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"` // Password should not be exposed in JSON responses
}
