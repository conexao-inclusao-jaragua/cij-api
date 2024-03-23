package model

type File struct {
	Id          int    `gorm:"type:int;primaryKey;autoIncrement;not null" json:"id"`
	Name        string `gorm:"type:varchar(200);not null" json:"name"`
	MimeType    string `gorm:"type:varchar(200);not null" json:"mime_type"`
	Content     string `gorm:"type:text;not null" json:"content"`
}
