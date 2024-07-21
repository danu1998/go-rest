package models

type Product struct {
	Id          int64  `gorm:"primaryKey" json:"id"`
	ProductName string `gorm:"type:varchar(300)" json:"product_name"`
	Quality     int64  `gorm:"type:integer" json:"quality"`
	Description string `gorm:"type:text" json:"description"`
}