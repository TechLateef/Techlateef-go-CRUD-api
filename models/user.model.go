package models

import (
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"uniqueIndex;type:varchar(255); not null" json:"username"`
	Password string `gorm:"type:varchar(255);not null" json:"password"`
}

// this func will save the user in our db

func (u *User) SaveUser() (*User, error) {

	var err error
	err = DB.Save(&u).Error
	{
		return &User{}, err
	}

}

// this fun will be running the hash proccess to the user's password

func (u *User) BeforeSave() error {

	// turn password into hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.MinCost)

	if err != nil {
		log.Println(err)
		panic("failed to hash password")
	}

	//remove spaces in username

	return string(hashedPassword)
}
