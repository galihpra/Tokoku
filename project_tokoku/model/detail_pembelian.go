package model

import (
	"time"

	"gorm.io/gorm"
)

type DetailPembelian struct {
	PembelianID string `gorm:"type:varchar(100);primaryKey"`
	ProductID   string `gorm:"type:varchar(20);primaryKey"`
	Sub_total   int
	Qty         int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
