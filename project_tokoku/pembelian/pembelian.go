package pembelian

import (
	"fmt"
	"project_tokoku/model"

	"gorm.io/gorm"
)

type PembelianSystem struct {
	DB *gorm.DB
}

func (ps *PembelianSystem) CreatePembelian(userID string) (model.Pembelian, bool) {
	var newPembelian = new(model.Pembelian)
	fmt.Print("Masukkan Nomor Invoice: ")
	fmt.Scanln(&newPembelian.No_invoice)
	fmt.Print("Masukkan Nomor HP Customer: ")
	fmt.Scanln(&newPembelian.CustomerID)
	newPembelian.UserID = userID

	err := ps.DB.Create(newPembelian).Error
	if err != nil {
		fmt.Println("input error:", err.Error())
		return model.Pembelian{}, false
	}

	return *newPembelian, true
}
