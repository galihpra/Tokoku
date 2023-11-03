package products

import (
	"bufio"
	"fmt"
	"os"
	"project_tokoku/model"
	"strings"

	"gorm.io/gorm"
)

type ProductSystem struct {
	DB *gorm.DB
}

func (ps *ProductSystem) CreateProduct(userID string) (model.Product, bool) {
	var newProduct = new(model.Product)
	fmt.Print("Masukkan Barcode Produk: ")
	fmt.Scanln(&newProduct.Barcode)

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan Nama Produk: ")
	name, _ := reader.ReadString('\n')
	newProduct.Nama = strings.TrimSpace(name)

	fmt.Print("Masukkan Harga Produk: ")
	fmt.Scanln(&newProduct.Harga)
	fmt.Print("Masukkan Stok Produk: ")
	fmt.Scanln(&newProduct.Stok)
	newProduct.UserID = userID

	err := ps.DB.Create(newProduct).Error
	if err != nil {
		fmt.Println("input error:", err.Error())
		return model.Product{}, false
	}

	return *newProduct, true
}

func (ps *ProductSystem) ReadProducts() ([]model.Product, bool) {
	var listProduk []model.Product

	err := ps.DB.Model(&model.Product{}).
		Select("products.*, users.nama as nama").
		Joins("JOIN users on products.user_id = users.username").
		Scan(&listProduk).
		Error

	if err != nil {
		fmt.Println("Error:", err.Error())
		return nil, false
	}

	return listProduk, true
}

func (ps *ProductSystem) UpdateInfoProduk(barcode string, produkUpdate model.Product) bool {
	var produk model.Product
	qry := ps.DB.Where("barcode = ?", barcode).First(&produk)
	if qry.Error != nil {
		fmt.Println("Produk tidak ditemukan")
		return false
	}

	produk.Nama = produkUpdate.Nama
	produk.Harga = produkUpdate.Harga
	produk.UserID = produkUpdate.UserID

	if err := ps.DB.Model(&produk).Updates(&produk).Error; err != nil {
		fmt.Println("Gagal mengupdate produk:", err.Error())
		return false
	}

	return true
}

func (ps *ProductSystem) GetProductsByID(Barcode []string) ([]model.Product, bool) {
	var listProduk []model.Product

	for i := 0; i < len(Barcode); i++ {
		var produk model.Product
		err := ps.DB.Where("barcode = ?", Barcode[i]).Find(&produk).Error

		if err != nil {
			fmt.Println("Error:", err.Error())
			return nil, false
		}

		listProduk = append(listProduk, produk)
	}

	return listProduk, true
}

func (ps *ProductSystem) UpdateStokProduk(barcode string, produkUpdate model.Product) bool {
	var produk model.Product
	qry := ps.DB.Where("barcode = ?", barcode).First(&produk)
	if qry.Error != nil {
		fmt.Println("Produk tidak ditemukan")
		return false
	}

	produk.Stok = produkUpdate.Stok

	if err := ps.DB.Model(&produk).Updates(&produk).Error; err != nil {
		fmt.Println("Gagal mengupdate produk:", err.Error())
		return false
	}

	return true
}

func (ps *ProductSystem) DeleteProduct(barcode string) bool {
	var produkData model.Product

	qry := ps.DB.Where("barcode = ?", barcode).First(&produkData)
	if qry.Error != nil {
		fmt.Println("Customer tidak ditemukan")
		return false
	}

	if err := ps.DB.Delete(&produkData).Error; err != nil {
		fmt.Println("Gagal menghapus produk: ", err.Error())
		return false
	}

	return true
}
