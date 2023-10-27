package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Barcode string
	Nama    string
	Harga   int
	Stok    int
}
