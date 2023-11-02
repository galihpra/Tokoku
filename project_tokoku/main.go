package main

import (
	"fmt"
	"project_tokoku/auth"
	"project_tokoku/config"
	"project_tokoku/customer"
	"project_tokoku/detailPembelian"
	"project_tokoku/model"
	"project_tokoku/pembelian"
	"project_tokoku/products"
	"project_tokoku/users"
	"time"
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
	var pembelian = pembelian.PembelianSystem{DB: db}
	var detailPembelian = detailPembelian.DetailPembelianSystem{DB: db}

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
					fmt.Println("5. Detail Transaksi")
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
							fmt.Println("5. Hapus Produk")
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
										fmt.Printf("Barcode : %s \nNama Editor : %s \nNama Barang : %s \nHarga : %d \nStok : %d \n", a.Barcode, a.UserNama, a.Nama, a.Harga, a.Stok)
										fmt.Println("================================")
										// fmt.Println(a)
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
						var menuTransaksi int
						var menuTransaksiActive bool = true
						for menuTransaksiActive {
							fmt.Println("Menu Transaksi:")
							fmt.Println("1. Lihat Daftar Produk")
							fmt.Println("2. Pilih Produk")
							fmt.Println("3. Ubah Informasi Transaksi")
							fmt.Println("4. Hapus Transaksi")
							fmt.Println("0. Kembali")
							fmt.Print("Pilih Menu: ")
							fmt.Scanln(&menuTransaksi)
							switch menuTransaksi {
							case 1:
								result, permit := products.ReadProducts()
								if permit {
									fmt.Println("====================================")
									fmt.Println(" Barcode    Produk     Harga    Stok")
									fmt.Println("------------------------------------")
									for _, a := range result {
										fmt.Println(a.Barcode, "  ", a.Nama, "  ", a.Harga, "  ", a.Stok)
									}
									fmt.Println("====================================")
								}
							case 2:
								var PilihBarang []string
								var Jumlah []int
								var barcode string
								var jml int

								for {
									fmt.Println("Masukkan Barcode Produk: ")
									fmt.Scanln(&barcode)
									fmt.Println("Masukkan Jumlah Produk: ")
									fmt.Scanln(&jml)

									PilihBarang = append(PilihBarang, barcode)
									Jumlah = append(Jumlah, jml)

									fmt.Print("Tambahkan Produk Lainnya? (y/n): ")
									var pilihan string
									fmt.Scanln(&pilihan)

									if pilihan != "y" {
										break
									}
								}

								fmt.Println(PilihBarang)
								fmt.Println(Jumlah)

							case 3:
							case 4:
							case 0:
								menuTransaksiActive = false
							}
						}
					case 5:
						var menuDetailTransaksi int
						var menuDetailTransaksiActive bool = true
						for menuDetailTransaksiActive {
							fmt.Println("Menu Detail Transaksi:")
							fmt.Println("1. Masukkan Produk")
							fmt.Println("2. Cetak Struk")
							fmt.Println("3. Ubah Detail Transaksi")
							fmt.Println("4. Hapus Detail Transaksi")
							fmt.Println("0. Kembali")
							fmt.Print("Pilih Menu: ")
							fmt.Scanln(&menuDetailTransaksi)
							switch menuDetailTransaksi {
							case 1:
								result, permit := detailPembelian.CreateDetailPembelian()
								if permit {
									fmt.Println(result)
								}
							case 2:
								jam := time.Now()
								r, permit := detailPembelian.ReadDetailPembelian()
								var total int
								if permit {
									fmt.Println("======================================")
									fmt.Println("Staff :", result.Nama)
									fmt.Println(jam.Format("2006-01-02"))
									fmt.Println("Nomor Invoice: ", r[0].PembelianID)
									fmt.Println("--------------------------------------")
									fmt.Println("Barang", "   ", "Jumlah", "     ", "Sub total")
									fmt.Println("--------------------------------------")
									for _, a := range r {
										fmt.Println(a.ProdukNama, "      ", a.Qty, "          ", a.Sub_total)
										total += a.Sub_total
									}
									fmt.Println("--------------------------------------")
									fmt.Println("Total                      ", total)
									fmt.Println("======================================")
								}
							case 3:
							case 4:
							case 0:
								menuDetailTransaksiActive = false
							}
						}
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
					fmt.Println("4. Detail Transaksi")
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
						var menuCustomer int
						var menuCustomerActive bool = true
						for menuCustomerActive {
							fmt.Println("Menu Customer:")
							fmt.Println("1. Tambahkan Customer")
							fmt.Println("2. Lihat Daftar Customer")
							fmt.Println("3. Ubah Informasi Customer")
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
							case 0:
								menuCustomerActive = false
							}
						}
					case 3:
						var menuTransaksi int
						var menuTransaksiActive bool = true
						for menuTransaksiActive {
							fmt.Println("Menu Transaksi:")
							fmt.Println("1. Buat Transaksi")
							fmt.Println("2. Lihat Daftar Transaksi")
							fmt.Println("3. Ubah Informasi Transaksi")
							fmt.Println("0. Kembali")
							fmt.Print("Pilih Menu: ")
							fmt.Scanln(&menuTransaksi)
							switch menuTransaksi {
							case 1:
								result, permit := pembelian.CreatePembelian(result.Username)
								if permit {
									fmt.Println(result)
								}
							case 2:
								result, permit := pembelian.ReadPembelian()
								if permit {
									for _, a := range result {
										fmt.Println(a)
									}
								}
							case 3:
							case 4:
							case 0:
								menuTransaksiActive = false
							}
						}
					case 4:
						var menuDetailTransaksi int
						var menuDetailTransaksiActive bool = true
						for menuDetailTransaksiActive {
							fmt.Println("Menu Detail Transaksi:")
							fmt.Println("1. Masukkan Produk")
							fmt.Println("2. Cetak Struk")
							fmt.Println("0. Kembali")
							fmt.Print("Pilih Menu: ")
							fmt.Scanln(&menuDetailTransaksi)
							switch menuDetailTransaksi {
							case 1:
								result, permit := detailPembelian.CreateDetailPembelian()
								if permit {
									fmt.Println(result)
								}
							case 2:
								jam := time.Now()
								r, permit := detailPembelian.ReadDetailPembelian()
								var total int
								if permit {
									fmt.Println("======================================")
									fmt.Println("Staff :", result.Nama)
									fmt.Println(jam.Format("2006-01-02"))
									fmt.Println("Nomor Invoice: ", r[0].PembelianID)
									fmt.Println("--------------------------------------")
									fmt.Println("Barang", "   ", "Jumlah", "     ", "Sub total")
									fmt.Println("--------------------------------------")
									for _, a := range r {
										fmt.Println(a.ProdukNama, "      ", a.Qty, "          ", a.Sub_total)
										total += a.Sub_total
									}
									fmt.Println("--------------------------------------")
									fmt.Println("Total                      ", total)
									fmt.Println("======================================")
								}
							case 0:
								menuDetailTransaksiActive = false
							}
						}
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
