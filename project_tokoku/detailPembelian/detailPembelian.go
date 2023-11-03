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

func buatInvoice() string {
	var tanggal = time.Now()
	return fmt.Sprintf("TKK-%d%d%d%d%d", tanggal.Year(), tanggal.Month(), tanggal.Day(), tanggal.Minute(), tanggal.Second())
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

		// Validasi stok cukup
		if product.Stok < newDetailPembelian.Qty {
			fmt.Println(">>>>>>>>STOK TIDAK CUKUP!!!<<<<<<<<<")
			return model.DetailPembelian{}, false
		}

		err := dps.DB.Create(newDetailPembelian).Error
		if err != nil {
			fmt.Println("input error:", err.Error())
			return model.DetailPembelian{}, false
		}

		product.Stok -= newDetailPembelian.Qty

		if product.Stok < 0 {
			product.Stok = 0
		}

		if err := dps.DB.Save(&product).Error; err != nil {
			fmt.Println("error mengurangi stok:", err.Error())
			return model.DetailPembelian{}, false
		}
	}

	return *newDetailPembelian, true
}

func (dps *DetailPembelianSystem) ReadDetailPembelian(invoice string) ([]model.DetailPembelian, bool) {
	var listDetail []model.DetailPembelian

	// var invoice string
	// fmt.Println("Masukkan Nomor Invoice: ")
	// fmt.Scanln(&invoice)

	err := dps.DB.Where("pembelian_id = ?", invoice).Model(&model.DetailPembelian{}).
		Select("detail_pembelians.pembelian_id, detail_pembelians.product_id,detail_pembelians.qty, detail_pembelians.sub_total, products.nama as nama").
		Joins("JOIN products on detail_pembelians.product_id = products.barcode").
		Scan(&listDetail).
		Error

	if err != nil {
		fmt.Println("Error:", err.Error())
		return nil, false
	}

	return listDetail, true
}

func (dps *DetailPembelianSystem) UpdateDetailPembelian(barcode, invoice string, detailPembelianUpdate model.DetailPembelian) bool {
	var details model.DetailPembelian
	qry := dps.DB.Where("product_id = ? AND pembelian_id = ?", barcode, invoice).First(&details)
	if qry.Error != nil {
		fmt.Println("Detail transaksi tidak ditemukan")
		return false
	}

	details.Qty = detailPembelianUpdate.Qty

	product := model.Product{}
	cari := dps.DB.Where("barcode = ?", barcode).Take(&product).Error
	if cari != nil {
		fmt.Println("input error:", cari.Error())
		return false
	}
	details.Sub_total = details.Qty * product.Harga

	if err := dps.DB.Model(&details).Updates(&details).Error; err != nil {
		fmt.Println("Gagal mengupdate detail transaksi: ", err.Error())
		return false
	}

	return true
}

func (dps *DetailPembelianSystem) DeleteDetail(Barcode, Invoice string) bool {
	var details model.DetailPembelian

	qry := dps.DB.Where("product_id = ? AND pembelian_id = ?", Barcode, Invoice).First(&details)
	if qry.Error != nil {
		fmt.Println("Detail produk tidak ditemukan")
		return false
	}

	if err := dps.DB.Delete(&details).Error; err != nil {
		fmt.Println("Gagal menghapus detail transaksi: ", err.Error())
		return false
	}

	return true
}
