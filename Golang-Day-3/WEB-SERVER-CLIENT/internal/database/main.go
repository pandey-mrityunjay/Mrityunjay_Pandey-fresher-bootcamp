package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	fmt.Println("Inside Database Main.go ----------------->")
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	DB = database
	fmt.Println("✅ Database connected successfully")

	err = DB.AutoMigrate(&Student{})
	if err != nil {
		log.Fatal("❌ Migration failed:", err)
	}
	log.Println("✅ Database connected and migration complete!")
	return DB
}
func GetDB() *gorm.DB {
	return DB
}
