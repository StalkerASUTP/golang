package user

import (
	"gorm.io/gorm"
)

const randomaizer int = 6

type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"index"`
	Password string
}
