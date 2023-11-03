package main

import (
	"bufio"
	"fmt"
	"os"
	"project_tokoku/auth"
	"project_tokoku/config"
	"project_tokoku/customer"
	"project_tokoku/detailPembelian"
	"project_tokoku/model"
	"project_tokoku/pembelian"
	"project_tokoku/products"
	"project_tokoku/users"
	"strings"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

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
										fmt.Printf("Nama : %s \nUsername : %s \nPassword : %s \n", a.Nama, a.Username, a.Password)
										fmt.Println("================================")
									}
								}
								var menuInUser int
								var menuInUserActive bool = true
								for menuInUserActive {
									fmt.Println("Menu User: ")
									fmt.Println("1. Ubah Data User")
									fmt.Println("2. Hapus User")
									fmt.Println("0. Kembali")
									fmt.Println("Pilih Menu: ")
									fmt.Scanln(&menuInUser)
									switch menuInUser {
									case 1:
										var username string
										var userUpdate model.User
										fmt.Print("Masukkan username: ")
										fmt.Scanln(&username)

										fmt.Print("Masukkan Nama: ")
										name, _ := reader.ReadString('\n')
										userUpdate.Nama = strings.TrimSpace(name)
										fmt.Print("Masukkan Password: ")
										fmt.Scanln(&userUpdate.Password)

										success := users.UpdateUser(username, userUpdate)

										if success {
											fmt.Println("Produk berhasil diubah")
										}
									case 2:
										var username string
										fmt.Print("Masukkan Username: ")
										fmt.Scanln(&username)
										sucess := users.DeleteUser(username)
										if sucess {
											fmt.Println("Data user berhasil dihapus")
										}
									case 0:
										menuInUserActive = false
									}
								}
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
									fmt.Println("====================================")
									fmt.Println(" Barcode    Produk     Harga    Stok    InputBy")
									fmt.Println("------------------------------------")
									for _, a := range result {
										fmt.Println(a.Barcode, "  ", a.Nama, "  ", a.Harga, "  ", a.Stok, "  ", a.UserNama)
									}
									fmt.Println("====================================")
								}
							case 3:
								hasil, permit := products.ReadProducts()
								if permit {
									fmt.Println("====================================")
									fmt.Println(" Barcode    Produk     Harga    Stok")
									fmt.Println("------------------------------------")
									for _, a := range hasil {
										fmt.Println(a.Barcode, "  ", a.Nama, "  ", a.Harga, "  ", a.Stok)
									}
									fmt.Println("====================================")
								}
								var barcode string
								var produkUpdate model.Product
								fmt.Print("Masukkan barcode produk: ")
								fmt.Scanln(&barcode)

								fmt.Print("Masukkan Nama Produk: ")
								name, _ := reader.ReadString('\n')
								produkUpdate.Nama = strings.TrimSpace(name)
								fmt.Print("Masukkan Harga Produk: ")
								fmt.Scanln(&produkUpdate.Harga)
								produkUpdate.UserID = result.Username

								success := products.UpdateInfoProduk(barcode, produkUpdate)

								if success {
									fmt.Println("Produk berhasil diubah")
								}

							case 4:
								hasil, permit := products.ReadProducts()
								if permit {
									fmt.Println("====================================")
									fmt.Println(" Barcode    Produk     Harga    Stok")
									fmt.Println("------------------------------------")
									for _, a := range hasil {
										fmt.Println(a.Barcode, "  ", a.Nama, "  ", a.Harga, "  ", a.Stok)
									}
									fmt.Println("====================================")
								}
								var barcode string
								var produkUpdate model.Product
								fmt.Print("Masukkan barcode produk: ")
								fmt.Scanln(&barcode)

								fmt.Print("Masukkan Stok Produk: ")
								fmt.Scanln(&produkUpdate.Stok)

								success := products.UpdateStokProduk(barcode, produkUpdate)

								if success {
									fmt.Println("Stok produk berhasil diubah")
								}
							case 5:
								hasil, permit := products.ReadProducts()
								if permit {
									fmt.Println("====================================")
									fmt.Println(" Barcode    Produk     Harga    Stok")
									fmt.Println("------------------------------------")
									for _, a := range hasil {
										fmt.Println(a.Barcode, "  ", a.Nama, "  ", a.Harga, "  ", a.Stok)
									}
									fmt.Println("====================================")
								}
								var barcode string
								fmt.Print("Masukkan barcode produk: ")
								fmt.Scanln(&barcode)
								sucess := products.DeleteProduct(barcode)
								if sucess {
									fmt.Println("Data produk berhasil dihapus")
								}
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
									fmt.Println("====================================")
									fmt.Println("   Barcode         Nama Customer")
									fmt.Println("------------------------------------")
									for _, a := range result {
										fmt.Println(a.Hp, "            ", a.Nama)
									}
									fmt.Println("====================================")
								}
							case 3:
								result, permit := customer.ReadCustomer()
								if permit {
									fmt.Println("====================================")
									fmt.Println("   Barcode         Nama Customer")
									fmt.Println("------------------------------------")
									for _, a := range result {
										fmt.Println(a.Hp, "            ", a.Nama)
									}
									fmt.Println("====================================")
								}
								var hp string
								var customerUpdate model.Customer
								fmt.Print("Masukkan Nomor HP: ")
								fmt.Scanln(&hp)

								fmt.Print("Masukkan Nama Customer: ")
								name, _ := reader.ReadString('\n')
								customerUpdate.Nama = strings.TrimSpace(name)

								success := customer.UpdateCustomer(hp, customerUpdate)

								if success {
									fmt.Println("Customer berhasil diubah")
								}
							case 4:
								result, permit := customer.ReadCustomer()
								if permit {
									fmt.Println("====================================")
									fmt.Println("   Barcode         Nama Customer")
									fmt.Println("------------------------------------")
									for _, a := range result {
										fmt.Println(a.Hp, "            ", a.Nama)
									}
									fmt.Println("====================================")
								}
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
							fmt.Println("3. Lihat Daftar Transaksi")
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

								listCustomer, permit := customer.ReadCustomer()
								if permit {
									fmt.Println("====================================")
									fmt.Println("   Barcode         Nama Customer")
									fmt.Println("------------------------------------")
									for _, a := range listCustomer {
										fmt.Println(a.Hp, "            ", a.Nama)
									}
									fmt.Println("====================================")
								}

								var HP string
								fmt.Print("Masukkan Nomor HP Customer: ")
								fmt.Scanln(&HP)

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
										hasil, permit := products.GetProductsByID(PilihBarang)
										if permit {
											var i int
											var total int
											fmt.Println("====================================")
											fmt.Println(" Produk    Jumlah     Sub Total")
											fmt.Println("------------------------------------")
											for _, a := range hasil {
												fmt.Println(a.Nama, "    ", Jumlah[i], "    ", a.Harga*Jumlah[i])
												total += a.Harga * Jumlah[i]
												i++
											}
											fmt.Println("              total : ", total)
											fmt.Println("====================================")

											var simpanTransaksi string
											fmt.Print("Buat Transaksi? (y/n): ")
											fmt.Scanln(&simpanTransaksi)
											if simpanTransaksi == "y" {
												pembelian.CreatePembelian(HP, result.Username, total)
												detailPembelian.CreateDetailPembelian(PilihBarang, Jumlah)
												PilihBarang = nil
												Jumlah = nil
												fmt.Println("*********Transaksi Selesai*********")

												//nota
												var tanggal = time.Now()
												r, permit := detailPembelian.ReadDetailPembelian(fmt.Sprintf("TKK-%d%d%d%d%d", tanggal.Year(), tanggal.Month(), tanggal.Day(), tanggal.Minute(), tanggal.Second()))
												var total int
												if permit {
													fmt.Println("======================================")
													fmt.Println("Staff :", result.Nama)
													fmt.Println(tanggal.Format("2006-01-02"))
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

												break
											} else if simpanTransaksi == "n" {
												PilihBarang = nil
												Jumlah = nil
												fmt.Println("*********Transaksi Dibatalkan*********")
												break
											}
										}
									}
								}
							case 3:
								result, permit := pembelian.ReadPembelian()
								if permit {
									fmt.Println("====================================")
									fmt.Println(" Invoice    Customer     Total")
									fmt.Println("------------------------------------")
									for _, a := range result {
										fmt.Println(a.No_invoice, "    ", a.CustomerID, "    ", a.Total)
									}
									fmt.Println("====================================")
								}
								var menuDaftarTransaksi int
								var menuDaftarTransaksiActive bool = true
								for menuDaftarTransaksiActive {
									fmt.Println("Menu Lihat Daftar Transaksi:")
									fmt.Println("1. Lihat Detail Transaksi")
									fmt.Println("2. Hapus Transaksi")
									fmt.Println("0. Kembali")
									fmt.Print("Pilih Menu: ")
									fmt.Scanln(&menuDaftarTransaksi)
									switch menuDaftarTransaksi {
									case 1:
										var invoice string
										fmt.Println("Masukkan Nomor Invoice: ")
										fmt.Scanln(&invoice)
										result, permit := detailPembelian.ReadDetailPembelian(invoice)
										if permit {
											fmt.Println("====================================")
											fmt.Println("Barang", "   ", "Jumlah", "     ", "Sub total")
											fmt.Println("--------------------------------------")
											for _, a := range result {
												fmt.Println(a.ProductID, a.ProdukNama, "      ", a.Qty, "          ", a.Sub_total)
											}
											fmt.Println("====================================")
										}
										var menuDetailTransaksi int
										var menuDetailTransaksiActive bool = true
										for menuDetailTransaksiActive {
											fmt.Println("Menu Detail Transaksi: ")
											fmt.Println("1. Ubah Detail Transaksi")
											fmt.Println("2. Hapus Produk Detail Transaksi")
											fmt.Println("0. Kembali")
											fmt.Print("Pilih Menu: ")
											fmt.Scanln(&menuDetailTransaksi)
											switch menuDetailTransaksi {
											case 1:
												var barcode string
												var detailTransaksiUpdate model.DetailPembelian
												fmt.Print("Masukkan Barcode: ")
												fmt.Scanln(&barcode)

												fmt.Print("Masukkan Jumlah Produk: ")
												fmt.Scanln(&detailTransaksiUpdate.Qty)

												success := detailPembelian.UpdateDetailPembelian(barcode, invoice, detailTransaksiUpdate)

												if success {
													fmt.Println("Detail transaksi berhasil diubah")
												}
											case 2:
												var Barcode string
												fmt.Print("Masukkan Barcode: ")
												fmt.Scanln(&Barcode)
												sucess := detailPembelian.DeleteDetail(Barcode, invoice)
												if sucess {
													fmt.Println("Data detail transaksi berhasil dihapus")
												}
											case 0:
												menuDetailTransaksiActive = false
											}
										}
									case 2:
										var invoice string
										fmt.Println("Masukkan Nomor Invoice: ")
										fmt.Scanln(&invoice)
										sucess := pembelian.DeletePembelian(invoice)
										if sucess {
											fmt.Println("Data transaksi berhasil dihapus")
										}
									case 0:
										menuDaftarTransaksiActive = false
									}
								}
							case 0:
								menuTransaksiActive = false
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
									fmt.Println("====================================")
									fmt.Println(" Barcode    Produk     Harga    Stok    InputBy")
									fmt.Println("------------------------------------")
									for _, a := range result {
										fmt.Println(a.Barcode, "  ", a.Nama, "  ", a.Harga, "  ", a.Stok, "  ", a.UserNama)
									}
									fmt.Println("====================================")
								}
							case 3:
								hasil, permit := products.ReadProducts()
								if permit {
									fmt.Println("====================================")
									fmt.Println(" Barcode    Produk     Harga    Stok")
									fmt.Println("------------------------------------")
									for _, a := range hasil {
										fmt.Println(a.Barcode, "  ", a.Nama, "  ", a.Harga, "  ", a.Stok)
									}
									fmt.Println("====================================")
								}
								var barcode string
								var produkUpdate model.Product
								fmt.Print("Masukkan barcode produk: ")
								fmt.Scanln(&barcode)

								fmt.Print("Masukkan Nama Produk: ")
								name, _ := reader.ReadString('\n')
								produkUpdate.Nama = strings.TrimSpace(name)
								fmt.Print("Masukkan Harga Produk: ")
								fmt.Scanln(&produkUpdate.Harga)
								produkUpdate.UserID = result.Username

								success := products.UpdateInfoProduk(barcode, produkUpdate)

								if success {
									fmt.Println("Produk berhasil diubah")
								}

							case 4:
								hasil, permit := products.ReadProducts()
								if permit {
									fmt.Println("====================================")
									fmt.Println(" Barcode    Produk     Harga    Stok")
									fmt.Println("------------------------------------")
									for _, a := range hasil {
										fmt.Println(a.Barcode, "  ", a.Nama, "  ", a.Harga, "  ", a.Stok)
									}
									fmt.Println("====================================")
								}
								var barcode string
								var produkUpdate model.Product
								fmt.Print("Masukkan barcode produk: ")
								fmt.Scanln(&barcode)

								fmt.Print("Masukkan Stok Produk: ")
								fmt.Scanln(&produkUpdate.Stok)

								success := products.UpdateStokProduk(barcode, produkUpdate)

								if success {
									fmt.Println("Stok produk berhasil diubah")
								}
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
									fmt.Println("====================================")
									fmt.Println("   Barcode         Nama Customer")
									fmt.Println("------------------------------------")
									for _, a := range result {
										fmt.Println(a.Hp, "            ", a.Nama)
									}
									fmt.Println("====================================")
								}
							case 3:
								result, permit := customer.ReadCustomer()
								if permit {
									fmt.Println("====================================")
									fmt.Println("   Barcode         Nama Customer")
									fmt.Println("------------------------------------")
									for _, a := range result {
										fmt.Println(a.Hp, "            ", a.Nama)
									}
									fmt.Println("====================================")
								}
								var hp string
								var customerUpdate model.Customer
								fmt.Print("Masukkan Nomor HP: ")
								fmt.Scanln(&hp)

								fmt.Print("Masukkan Nama Customer: ")
								name, _ := reader.ReadString('\n')
								customerUpdate.Nama = strings.TrimSpace(name)

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
							fmt.Println("1. Lihat Daftar Produk")
							fmt.Println("2. Pilih Produk")
							fmt.Println("3. Lihat Daftar Transaksi")
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

								listCustomer, permit := customer.ReadCustomer()
								if permit {
									fmt.Println("====================================")
									fmt.Println("   Barcode         Nama Customer")
									fmt.Println("------------------------------------")
									for _, a := range listCustomer {
										fmt.Println(a.Hp, "            ", a.Nama)
									}
									fmt.Println("====================================")
								}

								var HP string
								fmt.Print("Masukkan Nomor HP Customer: ")
								fmt.Scanln(&HP)

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
										hasil, permit := products.GetProductsByID(PilihBarang)
										if permit {
											var i int
											var total int
											fmt.Println("====================================")
											fmt.Println(" Produk    Jumlah     Sub Total")
											fmt.Println("------------------------------------")
											for _, a := range hasil {
												fmt.Println(a.Nama, "    ", Jumlah[i], "    ", a.Harga*Jumlah[i])
												total += a.Harga * Jumlah[i]
												i++
											}
											fmt.Println("              total : ", total)
											fmt.Println("====================================")

											var simpanTransaksi string
											fmt.Print("Buat Transaksi? (y/n): ")
											fmt.Scanln(&simpanTransaksi)
											if simpanTransaksi == "y" {
												pembelian.CreatePembelian(HP, result.Username, total)
												detailPembelian.CreateDetailPembelian(PilihBarang, Jumlah)
												PilihBarang = nil
												Jumlah = nil
												fmt.Println("*********Transaksi Selesai*********")

												//nota
												var tanggal = time.Now()
												r, permit := detailPembelian.ReadDetailPembelian(fmt.Sprintf("TKK-%d%d%d%d%d", tanggal.Year(), tanggal.Month(), tanggal.Day(), tanggal.Minute(), tanggal.Second()))
												var total int
												if permit {
													fmt.Println("======================================")
													fmt.Println("Staff :", result.Nama)
													fmt.Println(tanggal.Format("2006-01-02"))
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

												break
											} else if simpanTransaksi == "n" {
												PilihBarang = nil
												Jumlah = nil
												fmt.Println("*********Transaksi Dibatalkan*********")
												break
											}
										}
									}
								}
							case 3:
								result, permit := pembelian.ReadPembelian()
								if permit {
									fmt.Println("====================================")
									fmt.Println(" Invoice    Customer     Total")
									fmt.Println("------------------------------------")
									for _, a := range result {
										fmt.Println(a.No_invoice, "    ", a.CustomerID, "    ", a.Total)
									}
									fmt.Println("====================================")
								}
								var menuDaftarTransaksi int
								var menuDaftarTransaksiActive bool = true
								for menuDaftarTransaksiActive {
									fmt.Println("Menu Lihat Daftar Transaksi:")
									fmt.Println("1. Lihat Detail Transaksi")
									fmt.Println("0. Kembali")
									fmt.Print("Pilih Menu: ")
									fmt.Scanln(&menuDaftarTransaksi)
									switch menuDaftarTransaksi {
									case 1:
										var invoice string
										fmt.Println("Masukkan Nomor Invoice: ")
										fmt.Scanln(&invoice)
										result, permit := detailPembelian.ReadDetailPembelian(invoice)
										if permit {
											fmt.Println("====================================")
											fmt.Println("Barang", "   ", "Jumlah", "     ", "Sub total")
											fmt.Println("--------------------------------------")
											for _, a := range result {
												fmt.Println(a.ProductID, a.ProdukNama, "      ", a.Qty, "          ", a.Sub_total)
											}
											fmt.Println("====================================")
										}

									case 0:
										menuDaftarTransaksiActive = false
									}
								}
							case 0:
								menuTransaksiActive = false
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
