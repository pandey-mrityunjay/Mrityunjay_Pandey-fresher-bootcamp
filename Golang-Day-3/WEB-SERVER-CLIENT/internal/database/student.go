package database

import (
	"gorm.io/gorm"
)

// Student Model
type Student struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	DOB       string  `json:"dob"`
	Address   string  `json:"address"`
	Marks     float64 `json:"marks"`
}

// Migrate the Schema
func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(&Student{})
}
