package services

import (
	//"Web-Server/internal/database"
	"fmt"

	"gorm.io/gorm"
)

type NotesService struct {
	DB *gorm.DB // ✅ Store DB instance
}

type Note struct {
	Id   int
	Name string
}

// ✅ Constructor function to initialize NotesService
func NewNotesService(db *gorm.DB) *NotesService {
	return &NotesService{DB: db}
}

func (n *NotesService) GetNotesService() []Note {
	data := []Note{
		{Id: 1, Name: "Note 1"},
		{Id: 2, Name: "Note 2"},
	}
	return data
}

func (n *NotesService) CreateNotesService() Note {
	data := Note{
		Id: 3, Name: "Note 3",
	}
	return data
}

// ✅ Check database connection properly
func (n *NotesService) CheckDatabaseConnection() bool {
	fmt.Println("Inside Notes Service DB Connection")
	return n.DB != nil
}
