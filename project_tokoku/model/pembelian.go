package model

import "gorm.io/gorm"

type Pembelian struct {
	gorm.Model
	No_invoice string
	Qty        int
}
