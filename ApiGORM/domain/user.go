package domain

import (
	"log"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// User type contains specific fields for the user data structure.
type User struct {
	Base
	Name     string `json:"name" gorm:"type:varchar(255)"`
	Email    string `json:"email" gorm:"type:varchar(255);unique_index"`
	Password string `json:"-" gorm:"type:varchar(255)"`
	Token    string `json:"token" gorm:"type:varchar(255);unique_index"`
}

// NewUser funciton returns an user struct empty
func NewUser() *User {
	return &User{}
}

// Prepare function returns an erro if the data is not correct
func (user *User) Prepare() error {

	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Error during the password generation: %v", err)
		return err
	}
	user.Password = string(password)

	err = user.validate()
	if err != nil {
		log.Fatalf("Error during the user validatin: %v", err)
		return err
	}

	user.Token = uuid.NewString()
	return nil
}

func (user *User) validate() error {
	return nil
}
