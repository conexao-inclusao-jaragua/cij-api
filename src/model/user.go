package model

import "cij_api/src/enum"

type User struct {
	Id       int             `gorm:"type:int;primaryKey;autoIncrement;not null" json:"id"`
	Name     string          `gorm:"type:varchar(200);not null" json:"name"`
	Cpf      string          `gorm:"type:char(11);not null" json:"cpf"`
	Phone    string          `gorm:"type:varchar(255);not null" json:"phone"`
	Email    string          `gorm:"type:varchar(255);not null" json:"email"`
	Password string          `gorm:"type:varchar(255);not null" json:"password"`
	Gender   enum.GenderEnum `gorm:"type:varchar(255);not null" json:"gender"`
}
