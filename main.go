package main

import (
	"Log_rocket-Go_GIN_GORM/controllers"
	"Log_rocket-Go_GIN_GORM/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/books", controllers.FindBooks)         // GET all books endpoint
	r.POST("/books", controllers.CreateBook)       // POST endpoint to add books to the DB
	r.GET("/books/:id", controllers.FindBook)      // GET endpoint to get a book based on it's ID
	r.PATCH("/books/:id", controllers.UpdateBook)  // PATCH endpoint to update a speicific book based on it's ID
	r.DELETE("/books/:id", controllers.DeleteBook) // DELETE Endpoint to delete a specific book based on it's ID

	r.Run()
}
