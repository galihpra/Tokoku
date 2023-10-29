package main

import (
	"fmt"
	"project_tokoku/auth"
	"project_tokoku/config"
	"project_tokoku/model"
	"project_tokoku/users"
)

func main() {
	var inputMenu int
	db, err := config.InitDB()
	if err != nil {
		fmt.Println("Something happened", err.Error())
		return
	}

	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Customer{})
	db.AutoMigrate(&model.Pembelian{})
	db.AutoMigrate(&model.Product{})
	db.AutoMigrate(&model.DetailPembelian{})

	var auth = auth.AuthSystem{DB: db}
	var users = users.UserSystem{DB: db}

	for {
		fmt.Println("1. Login")
		fmt.Println("99. Exit")
		fmt.Print("Pilih Menu: ")
		fmt.Scanln(&inputMenu)
		switch inputMenu {
		case 1:
			var menuLogin int
			result, permit := auth.Login()
			if permit && result.Status == 1 {
				fmt.Println("Selamat datang di halaman admin", result.Nama)
				for permit {
					fmt.Println("1. Buat User Baru")
					fmt.Println("2. Logout")
					fmt.Println("3. Menu 3")
					fmt.Println("99. Exit")
					fmt.Print("Pilih Menu: ")
					fmt.Scanln(&menuLogin)
					switch menuLogin {
					case 1:
						result, permit := users.RegisterUser()
						if permit {
							fmt.Println(result)
						}
					case 2:
						permit = false
						fmt.Println("Anda sudah logout")
					case 3:
					case 99:
						fmt.Println("Thank You...")
						return
					}
				}
			} else {
				fmt.Println("Selamat datang", result.Nama)
			}

		case 99:
			fmt.Println("Thank You...")
			return
		default:
		}
	}

}
