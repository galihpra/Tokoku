package model

import (
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	Hp         string `gorm:"type:varchar(13);primaryKey"`
	Nama       string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	Pembelians []Pembelian
}
