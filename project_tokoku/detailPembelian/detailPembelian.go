package detailPembelian

import (
	"fmt"
	"project_tokoku/model"
	"time"

	"gorm.io/gorm"
)

type DetailPembelianSystem struct {
	DB *gorm.DB
}

var tanggal = time.Now()

func buatInvoice() string {
	return fmt.Sprintf("TKK--%d-%02d-%02d-%03d", tanggal.Year(), tanggal.Month(), tanggal.Day(), tanggal.Minute())
}

func (dps *DetailPembelianSystem) CreateDetailPembelian(Barcode []string, Jumlah []int) (model.DetailPembelian, bool) {
	var newDetailPembelian = new(model.DetailPembelian)

	for i := 0; i < len(Barcode); i++ {
		newDetailPembelian.PembelianID = buatInvoice()

		newDetailPembelian.ProductID = Barcode[i]
		newDetailPembelian.Qty = Jumlah[i]

		product := model.Product{}
		cari := dps.DB.Where("barcode = ?", newDetailPembelian.ProductID).Take(&product).Error
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
	}

	return *newDetailPembelian, true
}

func (dps *DetailPembelianSystem) ReadDetailPembelian() ([]model.DetailPembelian, bool) {
	var listDetail []model.DetailPembelian

	var invoice string
	fmt.Println("Masukkan Nomor Invoice: ")
	fmt.Scanln(&invoice)

	err := dps.DB.Where("pembelian_id = ?", invoice).Model(&model.DetailPembelian{}).
		Select("detail_pembelians.pembelian_id, detail_pembelians.qty, detail_pembelians.sub_total, products.nama as nama").
		Joins("JOIN products on detail_pembelians.product_id = products.barcode").
		Scan(&listDetail).
		Error

	if err != nil {
		fmt.Println("Error:", err.Error())
		return nil, false
	}

	return listDetail, true
}
