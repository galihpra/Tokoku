package main

import (
	"fmt"
	"project_tokoku/auth"
	"project_tokoku/config"
	"project_tokoku/customer"
	"project_tokoku/model"
	"project_tokoku/products"
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
	var products = products.ProductSystem{DB: db}
	var customer = customer.CustomerSystem{DB: db}

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
					fmt.Println("Menu Utama:")
					fmt.Println("1. User")
					fmt.Println("2. Produk")
					fmt.Println("3. Customer")
					fmt.Println("4. Transaksi")
					fmt.Println("0. Logout")
					fmt.Println("99. Exit")
					fmt.Print("Pilih Menu: ")
					fmt.Scanln(&menuLogin)
					switch menuLogin {
					case 1:
						var menuUser int
						var menuUserActive bool = true
						for menuUserActive {
							fmt.Println("Menu User:")
							fmt.Println("1. Tambahkan User")
							fmt.Println("2. Lihat Daftar User")
							fmt.Println("3. Ubah Informasi User")
							fmt.Println("4. Hapus User")
							fmt.Println("0. Kembali")
							fmt.Print("Pilih Menu: ")
							fmt.Scanln(&menuUser)
							switch menuUser {
							case 1:
								result, permit := users.RegisterUser()
								if permit {
									fmt.Println(result)
								}
							case 2:
								result, permit := users.ReadUser()
								if permit {
									for _, a := range result {
										fmt.Println(a)
									}
								}
							case 3:
							case 4:
							case 0:
								menuUserActive = false
							}
						}
					case 2:
						var menuProduk int
						var menuProdukActive bool = true
						for menuProdukActive {
							fmt.Println("Menu Produk:")
							fmt.Println("1. Tambahkan Produk")
							fmt.Println("2. Lihat Daftar Produk")
							fmt.Println("3. Ubah Informasi Produk")
							fmt.Println("4. Update Stok Produk")
							fmt.Println("5. Hapus User")
							fmt.Println("0. Kembali")
							fmt.Print("Pilih Menu: ")
							fmt.Scanln(&menuProduk)
							switch menuProduk {
							case 1:
								result, permit := products.CreateProduct(result.Username)
								if permit {
									fmt.Println(result)
								}
							case 2:
								result, permit := products.ReadProducts()
								if permit {
									for _, a := range result {
										fmt.Println(a)
									}
								}
							case 3:
								var barcode string
								var produkUpdate model.Product
								fmt.Print("Masukkan barcode produk: ")
								fmt.Scanln(&barcode)

								fmt.Print("Masukkan Nama Produk: ")
								fmt.Scanln(&produkUpdate.Nama)
								fmt.Print("Masukkan Harga Produk: ")
								fmt.Scanln(&produkUpdate.Harga)
								produkUpdate.UserID = result.Username

								success := products.UpdateInfoProduk(barcode, produkUpdate)

								if success {
									fmt.Println("Produk berhasil diubah")
								}

							case 4:
							case 5:
							case 0:
								menuProdukActive = false
							}
						}
					case 3:
						var menuCustomer int
						var menuCustomerActive bool = true
						for menuCustomerActive {
							fmt.Println("Menu Customer:")
							fmt.Println("1. Tambahkan Customer")
							fmt.Println("2. Lihat Daftar Customer")
							fmt.Println("3. Ubah Informasi Customer")
							fmt.Println("4. Hapus Customer")
							fmt.Println("0. Kembali")
							fmt.Print("Pilih Menu: ")
							fmt.Scanln(&menuCustomer)
							switch menuCustomer {
							case 1:
								result, permit := customer.CreateCustomer()
								if permit {
									fmt.Println(result)
								}
							case 2:
								result, permit := customer.ReadCustomer()
								if permit {
									for _, a := range result {
										fmt.Println(a)
									}
								}
							case 3:
								var hp string
								var customerUpdate model.Customer
								fmt.Print("Masukkan Nomor HP: ")
								fmt.Scanln(&hp)

								fmt.Print("Masukkan Nama Customer: ")
								fmt.Scanln(&customerUpdate.Nama)

								success := customer.UpdateCustomer(hp, customerUpdate)

								if success {
									fmt.Println("Customer berhasil diubah")
								}
							case 4:
								var customerID string
								fmt.Print("Masukkan Nomor HP: ")
								fmt.Scanln(&customerID)
								sucess := customer.DeleteCustomer(customerID)
								if sucess {
									fmt.Println("Data customer berhasil dihapus")
								}
							case 0:
								menuCustomerActive = false
							}
						}
					case 4:
					case 0:
						permit = false
						fmt.Println("Anda sudah logout")
					case 99:
						fmt.Println("Thank You...")
						return
					}
				}
			} else if permit && result.Status == 2 {
				fmt.Println("Selamat datang", result.Nama)
				for permit {
					fmt.Println("Menu Utama:")
					fmt.Println("1. Produk")
					fmt.Println("2. Customer")
					fmt.Println("3. Transaksi")
					fmt.Println("0. Logout")
					fmt.Println("99. Exit")
					fmt.Print("Pilih Menu: ")
					fmt.Scanln(&menuLogin)
					switch menuLogin {
					case 1:
						var menuProduk int
						var menuProdukActive bool = true
						for menuProdukActive {
							fmt.Println("Menu Produk:")
							fmt.Println("1. Tambahkan Produk")
							fmt.Println("2. Lihat Daftar Produk")
							fmt.Println("3. Ubah Informasi Produk")
							fmt.Println("4. Update Stok Produk")
							fmt.Println("5. Hapus User")
							fmt.Println("0. Kembali")
							fmt.Print("Pilih Menu: ")
							fmt.Scanln(&menuProduk)
							switch menuProduk {
							case 1:
								result, permit := products.CreateProduct(result.Username)
								if permit {
									fmt.Println(result)
								}
							case 2:
								result, permit := products.ReadProducts()
								if permit {
									for _, a := range result {
										fmt.Println(a)
									}
								}
							case 3:
							case 4:
							case 5:
							case 0:
								menuProdukActive = false
							}
						}
					case 2:
					case 3:
					case 0:
						permit = false
						fmt.Println("Anda sudah logout")
					case 99:
						fmt.Println("Thank You...")
						return
					}
				}
			}

		case 99:
			fmt.Println("Thank You...")
			return
		default:
		}
	}

}
