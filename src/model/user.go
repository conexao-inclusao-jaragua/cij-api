package model

import (
	"cij_api/src/enum"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       int             `gorm:"type:int;primaryKey;autoIncrement;not null" json:"id"`
	Name     string          `gorm:"type:varchar(200);not null" json:"name"`
	Cpf      string          `gorm:"type:char(11);not null;unique" json:"cpf"`
	Phone    string          `gorm:"type:char(13);not null" json:"phone"`
	Email    string          `gorm:"type:varchar(255);not null;unique" json:"email"`
	Password string          `gorm:"type:varchar(255);not null" json:"password"`
	IsAdmin  bool            `gorm:"type:tinyint(1);not null;default:0"`
	Gender   enum.GenderEnum `gorm:"type:char(6);not null" json:"gender"`
}

type UserResponse struct {
	Id     int             `json:"id"`
	Name   string          `json:"name"`
	Cpf    string          `json:"cpf"`
	Phone  string          `json:"phone"`
	Email  string          `json:"email"`
	Gender enum.GenderEnum `json:"gender"`
}

func (u *User) ValidatePassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)) == nil
}

func (u *User) ToResponse() UserResponse {
	return UserResponse{
		Id:     u.Id,
		Name:   u.Name,
		Cpf:    u.Cpf,
		Phone:  u.Phone,
		Email:  u.Email,
		Gender: u.Gender,
	}
}
