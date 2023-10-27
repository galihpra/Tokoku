package main

import (
	"fmt"
	"project_tokoku/config"
	"project_tokoku/model"
)

func main() {
	var inputMenu int
	db, err := config.InitDB()
	if err != nil {
		fmt.Println("Something happened", err.Error())
		return
	}

	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Pembelian{})
	db.AutoMigrate(&model.Customer{})
	db.AutoMigrate(&model.Product{})

	for {
		fmt.Println("1. Login")
		fmt.Println("99. Exit")
		fmt.Print("Pilih Menu: ")
		fmt.Scanln(&inputMenu)
		switch inputMenu {
		case 1:
		case 99:
			fmt.Println("Thank You...")
			return
		default:
		}
	}

}
