package models

//import "gorm.io/gorm"

//Product model represents the product table , containing all the info required in product details

type Product struct {
	ID          string  `gorm:"primaryKey" json:"id"`
	ProductName string  `gorm:"not null" json:"product_name"`
	Price       float64 `gorm:"not null" json:"price"`
	Quantity    int     `gorm:"not null" json:"quantity"`
}
