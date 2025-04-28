package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Email    string `gorm:"unique"`
	Password string
}

type RegisterInput struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
