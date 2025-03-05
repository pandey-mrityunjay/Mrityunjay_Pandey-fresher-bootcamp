package controllers

import (
	"Web-Server/services"

	"github.com/gin-gonic/gin"
)

type NotesController struct {
	NoteService *services.NotesService // Use pointer instead of value
}

func (n *NotesController) InitNotesControllerRoutes(router *gin.Engine) {
	notes := router.Group("/notes")
	notes.GET("/", n.GetNotes())
	notes.POST("/", n.CreateNotes())
	notes.GET("/ping", n.TestDbConnection())
}

func (n *NotesController) GetNotes() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"notes": n.NoteService.GetNotesService(),
		})
	}
}

func (n *NotesController) CreateNotes() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"notes": n.NoteService.CreateNotesService(),
		})
	}
}

// TestDbConnection function
func (n *NotesController) TestDbConnection() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"database": n.NoteService.CheckDatabaseConnection(),
		})
	}
}
