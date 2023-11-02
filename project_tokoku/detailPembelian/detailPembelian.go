package detailPembelian

import (
	"fmt"
	"project_tokoku/model"

	"gorm.io/gorm"
)

type DetailPembelianSystem struct {
	DB *gorm.DB
}

func (dps *DetailPembelianSystem) CreateDetailPembelian() (model.DetailPembelian, bool) {
	var newDetailPembelian = new(model.DetailPembelian)
	fmt.Print("Masukkan Nomor Invoice: ")
	fmt.Scanln(&newDetailPembelian.PembelianID)
	fmt.Print("Masukkan Barcode: ")
	fmt.Scanln(&newDetailPembelian.ProductID)
	fmt.Print("Masukkan Jumlah Barang: ")
	fmt.Scanln(&newDetailPembelian.Qty)

	product := model.Product{}
	cari := dps.DB.Where("barcode = ?", newDetailPembelian.ProductID).First(&product).Error
	if cari != nil {
		fmt.Println("input error:", cari.Error())
		return model.DetailPembelian{}, false
	}
	subTotal := product.Harga * newDetailPembelian.Qty
	newDetailPembelian.Sub_total = subTotal

	err := dps.DB.Create(newDetailPembelian).Error
	if err != nil {
		fmt.Println("input error:", err.Error())
		return model.DetailPembelian{}, false
	}

	return *newDetailPembelian, true
}
