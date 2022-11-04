package models

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	CustomModel
	Username  string `json:"username" gorm:"unique"`
	Email     string `json:"email" gorm:"unique"`
	Password  string `json:"password"`
	Privilege string `json:"privilege"`
}

func (user *User) HashPassword() error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return err
	}

	user.Password = string(bytes)
	return nil

}

func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(user.Password))
	if err != nil {
		return err
	}
	return nil
}