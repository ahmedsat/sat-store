package models

import (
	"golang.org/x/crypto/bcrypt"
)

// type privileges string

// const (
// 	USER  privileges = "USER"
// 	ADMEN privileges = "ADMEN"
// )

type User struct {
	CustomModel

	Username string `json:"username" gorm:"unique"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
	Phone    string `json:"Phone"`
	Address  string `json:"address"`

	Privileges string `sql:"type:ENUM('USER', 'ADMIN')" json:"privileges"`
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
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}

// func (self *privileges) Scan(value interface{}) error {
//     *self = privileges(value.([]byte))
//     return nil
// }

// func (self privileges) Value() (driver.Value, error) {
//     return string(self), nil
// }
