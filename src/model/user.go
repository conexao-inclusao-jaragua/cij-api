package model

import (
	"cij_api/src/enum"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       int             `gorm:"type:int;primaryKey;autoIncrement;not null" json:"-"`
	Name     string          `gorm:"type:varchar(200);not null" json:"name"`
	Cpf      string          `gorm:"type:char(11);not null;unique" json:"cpf"`
	Phone    string          `gorm:"type:char(13);not null" json:"phone"`
	Email    string          `gorm:"type:varchar(255);not null;unique" json:"email"`
	Password string          `gorm:"type:varchar(255);not null" json:"-"`
	Gender   enum.GenderEnum `gorm:"type:char(6);not null" json:"gender"`
}

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *User) ValidatePassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)) == nil
}
