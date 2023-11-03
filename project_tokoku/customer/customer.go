package customer

import (
	"bufio"
	"fmt"
	"os"
	"project_tokoku/model"
	"strings"

	"gorm.io/gorm"
)

type CustomerSystem struct {
	DB *gorm.DB
}

func (cs *CustomerSystem) CreateCustomer() (model.Customer, bool) {
	var newCustomer = new(model.Customer)
	fmt.Print("Masukkan Nomor HP: ")
	fmt.Scanln(&newCustomer.Hp)

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan Nama Produk: ")
	name, _ := reader.ReadString('\n')
	newCustomer.Nama = strings.TrimSpace(name)

	err := cs.DB.Create(newCustomer).Error
	if err != nil {
		fmt.Println("input error:", err.Error())
		return model.Customer{}, false
	}

	return *newCustomer, true
}

func (cs *CustomerSystem) ReadCustomer() ([]model.Customer, bool) {
	var customerList []model.Customer

	qry := cs.DB.Find(&customerList)
	err := qry.Error

	if err != nil {
		fmt.Println("Error read data table:", err.Error())
		return nil, false
	}

	return customerList, true
}

func (cs *CustomerSystem) DeleteCustomer(customerID string) bool {
	var customer model.Customer

	qry := cs.DB.Where("hp = ?", customerID).First(&customer)
	if qry.Error != nil {
		fmt.Println("Customer tidak ditemukan")
		return false
	}

	if err := cs.DB.Delete(&customer).Error; err != nil {
		fmt.Println("Gagal menghapus Customer: ", err.Error())
		return false
	}

	return true
}

func (cs *CustomerSystem) UpdateCustomer(hp string, customerUpdate model.Customer) bool {
	var customer model.Customer
	qry := cs.DB.Where("hp = ?", hp).First(&customer)
	if qry.Error != nil {
		fmt.Println("Customer tidak ditemukan")
		return false
	}

	customer.Nama = customerUpdate.Nama

	if err := cs.DB.Model(&customer).Updates(&customer).Error; err != nil {
		fmt.Println("Gagal mengupdate customer: ", err.Error())
		return false
	}

	return true
}
