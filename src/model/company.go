package model

import (
	"golang.org/x/crypto/bcrypt"
)

type Company struct {
	Id       int    `gorm:"type:int;primaryKey;autoIncrement;not null" json:"id"`
	Name     string `gorm:"type:varchar(200);not null" json:"name"`
	Cnpj     string `gorm:"type:char(13);not null;unique" json:"cnpj"`
	Phone    string `gorm:"type:char(13);not null" json:"phone"`
	Email    string `gorm:"type:varchar(255);not null;unique" json:"email"`
	Password string `gorm:"type:varchar(255);not null" json:"password"`
}

type CompanyResponse struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Cnpj  string `json:"cnpj"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}

func (c *Company) ValidatePassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(c.Password), []byte(password)) == nil
}

func (c *Company) ToResponse() CompanyResponse {
	return CompanyResponse{
		Id:    c.Id,
		Name:  c.Name,
		Cnpj:  c.Cnpj,
		Phone: c.Phone,
		Email: c.Email,
	}
}
