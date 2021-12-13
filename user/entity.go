package user

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name           string
	Occupation     string
	Email          string
	PasswordHash   string
	AvatarFileName string
	Role           string
}
