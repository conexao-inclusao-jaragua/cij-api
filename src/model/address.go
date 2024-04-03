package model

import (
	"time"

	"gorm.io/gorm"
)

type Address struct {
	*gorm.Model
	CreatedAt    time.Time `gorm:"<-:create"`
	Id           int       `gorm:"type:int;primaryKey;autoIncrement;not null" json:"id"`
	Street       string    `gorm:"type:varchar(200);not null" json:"street"`
	Number       string    `gorm:"type:varchar(200);not null" json:"number"`
	Neighborhood string    `gorm:"type:varchar(200);not null" json:"neighborhood"`
	City         string    `gorm:"type:varchar(200);not null" json:"city"`
	State        string    `gorm:"type:char(2);not null" json:"state"`
	Country      string    `gorm:"type:varchar(200);not null" json:"country"`
	ZipCode      string    `gorm:"type:char(8);not null" json:"zip_code"`
	Complement   string    `gorm:"type:varchar(200);" json:"complement"`
}

type AddressRequest struct {
	Street       string `json:"street"`
	Number       string `json:"number"`
	Neighborhood string `json:"neighborhood"`
	City         string `json:"city"`
	State        string `json:"state"`
	Country      string `json:"country"`
	ZipCode      string `json:"zip_code"`
	Complement   string `json:"complement"`
}

type AddressResponse struct {
	Id           int    `json:"id"`
	Street       string `json:"street"`
	Number       string `json:"number"`
	Neighborhood string `json:"neighborhood"`
	City         string `json:"city"`
	State        string `json:"state"`
	Country      string `json:"country"`
	ZipCode      string `json:"zip_code"`
	Complement   string `json:"complement"`
}

func (a *Address) ToResponse() AddressResponse {
	return AddressResponse{
		Id:           a.Id,
		Street:       a.Street,
		Number:       a.Number,
		Neighborhood: a.Neighborhood,
		City:         a.City,
		State:        a.State,
		Country:      a.Country,
		ZipCode:      a.ZipCode,
		Complement:   a.Complement,
	}
}

func (a *AddressRequest) ToModel() Address {
	return Address{
		Street:       a.Street,
		Number:       a.Number,
		Neighborhood: a.Neighborhood,
		City:         a.City,
		State:        a.State,
		Country:      a.Country,
		ZipCode:      a.ZipCode,
		Complement:   a.Complement,
	}
}
