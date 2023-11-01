package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Username   string `gorm:"type:varchar(55);primaryKey"`
	Nama       string
	Status     int
	Password   string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	Pembelians []Pembelian
	Products   []Product
}
