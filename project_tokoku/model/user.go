package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string
	Nama     string
	Status   string
	Password string
}
