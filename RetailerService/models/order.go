package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	ID         string         `gorm:"primaryKey" json:"id"`
	CustomerID string         `gorm:"not null" json:"customer_id"`
	ProductID  string         `gorm:"not null" json:"product_id"`
	Quantity   int            `gorm:"not null" json:"quantity"`
	Status     string         `gorm:"default:order placed" json:"status"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
}
