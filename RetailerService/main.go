package main

import (
	"log"
	"retailer-service/config"
	"retailer-service/models"
)

func main() {

	//Initialize database connection
	config.InitDB()

	//Run migrations to create tables
	err := config.DB.AutoMigrate(&models.Product{}, &models.Order{}, &models.Customer{})

	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	log.Println("Database migration completed successfully!!!!!!")
}
