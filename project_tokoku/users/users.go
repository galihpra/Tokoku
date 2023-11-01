package users

import (
	"fmt"
	"project_tokoku/model"

	"gorm.io/gorm"
)

type UserSystem struct {
	DB *gorm.DB
}

func (us *UserSystem) RegisterUser() (model.User, bool) {
	var newUser = new(model.User)
	fmt.Print("Masukkan Username: ")
	fmt.Scanln(&newUser.Username)
	fmt.Print("Masukkan Nama: ")
	fmt.Scanln(&newUser.Nama)
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
