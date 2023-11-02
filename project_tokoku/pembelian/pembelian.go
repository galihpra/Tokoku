package pembelian

import (
	"fmt"
	"project_tokoku/model"
	"time"

	"gorm.io/gorm"
)

type PembelianSystem struct {
	DB *gorm.DB
}

var tanggal = time.Now()

func buatInvoice() string {
	return fmt.Sprintf("TKK--%d-%02d-%02d-%03d", tanggal.Year(), tanggal.Month(), tanggal.Day(), tanggal.Minute())
}

func (ps *PembelianSystem) CreatePembelian(HP, userID string, total int) (model.Pembelian, bool) {
	var newPembelian = new(model.Pembelian)
	newPembelian.CustomerID = HP
	newPembelian.No_invoice = buatInvoice()
	newPembelian.UserID = userID
	newPembelian.Total = total

	err := ps.DB.Create(newPembelian).Error
	if err != nil {
		fmt.Println("input error:", err.Error())
		return model.Pembelian{}, false
	}

	return *newPembelian, true
}

func (ps *PembelianSystem) ReadPembelian() ([]model.Pembelian, bool) {
	var pembelianList []model.Pembelian

	qry := ps.DB.Find(&pembelianList)
	err := qry.Error

	if err != nil {
		fmt.Println("Error read data table:", err.Error())
		return nil, false
	}

	return pembelianList, true
}
