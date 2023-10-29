package model

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	Barcode          string `gorm:"type:varchar(20);primaryKey"`
	UserID           string `gorm:"type:varchar(55);foreignKey:Username"`
	Nama             string
	Harga            int
	Stok             int
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt `gorm:"index"`
	DetailPembelians []DetailPembelian
}
