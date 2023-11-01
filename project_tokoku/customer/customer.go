package customer

import (
	"fmt"
	"project_tokoku/model"

	"gorm.io/gorm"
)

type CustomerSystem struct {
	DB *gorm.DB
}

func (cs *CustomerSystem) CreateCustomer() (model.Customer, bool) {
	var newCustomer = new(model.Customer)
	fmt.Print("Masukkan Nomor HP: ")
	fmt.Scanln(&newCustomer.Hp)
	fmt.Print("Masukkan Nama: ")
	fmt.Scanln(&newCustomer.Nama)

	err := cs.DB.Create(newCustomer).Error
	if err != nil {
		fmt.Println("input error:", err.Error())
		return model.Customer{}, false
	}

	return *newCustomer, true
}
