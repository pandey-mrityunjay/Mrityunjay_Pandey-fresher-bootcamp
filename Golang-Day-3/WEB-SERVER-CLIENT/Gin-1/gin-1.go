package main

import (
	"Web-Server/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default() //Defining a router using gin framework
/* 	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK // inbuilt status return,
		gin.H{
			//gin.H : Inorder to a json message towards the client-side
			"message": "ping",
		})
	})

	router.GET("/me/:id/:newId", func(c *gin.Context) {
		var id = c.Param("id")       //way to extract the param from the req
		var newId = c.Param("newId") //way to extract the param from the req

		c.JSON(http.StatusOK, gin.H{
			"user_id":     id,
			"user_new_id": newId,
		})
	})

	router.POST("/me", func(c *gin.Context) {
		//Email and Password
		type MeRequest struct {
			Email    string `json:"email" binding:"required"` //binding makes this req field necesaary and it would throw an error if the requirements are not met
			Password string `json:password`
		}

		var meRequest MeRequest
		if err := c.BindJSON(&meRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"email":    meRequest.Email,
			"password": meRequest.Password,
		})
	})
	router.PUT("/me", func(c *gin.Context) {
		//Email and Password
		type MeRequest struct {
			Email    string `json:"email" binding:"required"` //binding makes this req field necesaary and it would throw an error if the requirements are not met
			Password string `json:password`
		}

		var meRequest MeRequest
		if err := c.BindJSON(&meRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"email":    meRequest.Email,
			"password": meRequest.Password,
		})
	})
	router.PATCH("/me", func(c *gin.Context) {
		//Email and Password
		type MeRequest struct {
			Email    string `json:"email" binding:"required"` //binding makes this req field necesaary and it would throw an error if the requirements are not met
			Password string `json:password`
		}

		var meRequest MeRequest
		if err := c.BindJSON(&meRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"email":    meRequest.Email,
			"password": meRequest.Password,
		})
	})

	router.DELETE("/me/:id", func(c *gin.Context) {
		id := c.Param("id")

		c.JSON(http.StatusOK, gin.H{
			"id":      id,
			"message": "deleted!!!",
		})
	}) */
	
	notesController:= &controllers.NotesController{}

	notesController.InitNotesControllerRoutes(router)
	
	
	router.Run(":8000")

}




//Controller redirects the request to the correct services, asks the services wait for its response -Middleware-Services that performs all the querying updation..etc etc all the tasks