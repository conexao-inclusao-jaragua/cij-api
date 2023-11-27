package model

import (
	"gorm.io/gorm"
)

type Disability struct {
	*gorm.Model
	Id          int    `gorm:"type:int;primaryKey;autoIncrement;not null" json:"id"`
	Category    string `gorm:"type:varchar(200);not null" json:"category"`
	Description string `gorm:"type:varchar(200);not null" json:"description"`
}

type DisabilityRequest struct {
	Category    string `json:"category"`
	Description string `json:"description"`
}

type DisabilityResponse struct {
	Id          int    `json:"id"`
	Category    string `json:"category"`
	Description string `json:"description"`
}

func (d *Disability) ToResponse() DisabilityResponse {
	return DisabilityResponse{
		Id:          d.Id,
		Category:    d.Category,
		Description: d.Description,
	}
}

func (d *DisabilityRequest) ToModel() Disability {
	return Disability{
		Category:    d.Category,
		Description: d.Description,
	}
}
