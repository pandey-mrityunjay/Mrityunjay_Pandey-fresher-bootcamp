package controllers

import (
	"Web-Server/internal/database"
	"Web-Server/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type StudentController struct {
	StudentService *services.StudentService
}

func NewStudentController(studentService *services.StudentService) *StudentController {
	return &StudentController{StudentService: studentService}
}

func (s *StudentController) RegisterRoutes(router *gin.Engine) {
	students := router.Group("/students")
	students.POST("/", s.CreateStudent)
	students.GET("/", s.GetStudents)
	students.GET("/:id", s.GetStudentByID)
	students.PUT("/:id", s.UpdateStudent)
	students.DELETE("/:id", s.DeleteStudent)
}

// Create Student
func (s *StudentController) CreateStudent(c *gin.Context) {
	var student database.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	if err := s.StudentService.CreateStudent(&student); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create studenrt"})
		return
	}
	c.JSON(http.StatusCreated, student)
}

// Get all students
func (s *StudentController) GetStudents(c *gin.Context) {
	students, err := s.StudentService.GetStudents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve students"})
		return
	}
	c.JSON(http.StatusOK, students)
}

// Get Student by ID
func (s *StudentController) GetStudentByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	student, err := s.StudentService.GetStudentByID(int(id))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Student not found"})
		return
	}
	c.JSON(http.StatusOK, student)
}

// Update Student
func (s *StudentController) UpdateStudent(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var student database.Student

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	if err := s.StudentService.UpdateStudent(int(id), &student); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update student"})
		return
	}
	c.JSON(http.StatusOK, student)
}

//Delete Student

func (s *StudentController) DeleteStudent(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	if err := s.StudentService.DeleteStudent(int(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete student"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Student deleted successfully"})
}
