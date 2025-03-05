package services

import (
	"Web-Server/internal/database"

	"gorm.io/gorm"
)

type StudentService struct {
	DB *gorm.DB
}

func NewStudentService(db *gorm.DB) *StudentService {
	return &StudentService{DB: db}
}

// Create Student
func (s *StudentService) CreateStudent(student *database.Student) error {
	return s.DB.Create(student).Error
}

// Get All Students
func (s *StudentService) GetStudents() ([]database.Student, error) {
	var students []database.Student
	err := s.DB.Find(&students).Error
	return students, err
}

// Get Student by ID
func (s *StudentService) GetStudentByID(id int) (*database.Student, error) {
	var student database.Student
	err := s.DB.First(&student, id).Error
	return &student, err
}

// Update Student
func (s *StudentService) UpdateStudent(id int, student *database.Student) error {
	var existingStudent database.Student

	// Check if student exists
	if err := s.DB.First(&existingStudent, id).Error; err != nil {
		return err // Student not found, return error
	}

	// Update only provided fields
	return s.DB.Model(&existingStudent).Updates(student).Error
}

// Delete Student
func (s *StudentService) DeleteStudent(id int) error {
	return s.DB.Delete(&database.Student{}, id).Error
}
