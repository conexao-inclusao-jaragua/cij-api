package model

type Role struct {
	Id   int    `gorm:"type:int;primaryKey;autoIncrement;not null" json:"id"`
	Name string `gorm:"type:varchar(200);not null" json:"name"`
}
