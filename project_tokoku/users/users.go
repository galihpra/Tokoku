package users

import (
	"bufio"
	"fmt"
	"os"
	"project_tokoku/model"
	"strings"

	"gorm.io/gorm"
)

type UserSystem struct {
	DB *gorm.DB
}

func (us *UserSystem) RegisterUser() (model.User, bool) {
	var newUser = new(model.User)
	fmt.Print("Masukkan Username: ")
	fmt.Scanln(&newUser.Username)

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan Nama: ")
	name, _ := reader.ReadString('\n')
	newUser.Nama = strings.TrimSpace(name)

	fmt.Print("Masukkan Password")
	fmt.Scanln(&newUser.Password)
	newUser.Status = 2

	err := us.DB.Create(newUser).Error
	if err != nil {
		fmt.Println("input error:", err.Error())
		return model.User{}, false
	}

	return *newUser, true
}

func (us *UserSystem) ReadUser() ([]model.User, bool) {
	var userList []model.User

	qry := us.DB.Find(&userList)
	err := qry.Error

	if err != nil {
		fmt.Println("Error read data table:", err.Error())
		return nil, false
	}

	return userList, true
}

func (us *UserSystem) UpdateUser(username string, userUpdate model.User) bool {
	var userData model.User
	qry := us.DB.Where("username = ?", username).First(&userData)
	if qry.Error != nil {
		fmt.Println("User tidak ditemukan")
		return false
	}

	userData.Nama = userUpdate.Nama
	userData.Password = userUpdate.Password

	if err := us.DB.Model(&userData).Updates(&userData).Error; err != nil {
		fmt.Println("Gagal mengupdate user:", err.Error())
		return false
	}

	return true
}

func (us *UserSystem) DeleteUser(username string) bool {
	var userData model.User

	qry := us.DB.Where("username = ?", username).First(&userData)
	if qry.Error != nil {
		fmt.Println("Customer tidak ditemukan")
		return false
	}

	if err := us.DB.Delete(&userData).Error; err != nil {
		fmt.Println("Gagal menghapus User: ", err.Error())
		return false
	}

	return true
}
