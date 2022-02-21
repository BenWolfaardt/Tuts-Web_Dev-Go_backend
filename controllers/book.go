package controllers

import (
	"Log_rocket-Go_GIN_GORM/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateBookInput struct {
	Title  string `json:"title" bindings:"required"`
	Author string `json:"author" binding:"required"`
}

type UpdatedBookInput struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

// GET /books endpoint
// Get all books
func FindBooks(c *gin.Context) {
	var books []models.Book
	models.DB.Find(&books)

	c.JSON(http.StatusOK, gin.H{"data": books})
}

// POST /books
// Create a new book in our DB
func CreateBook(c *gin.Context) {
	var input CreateBookInput

	// Validate the request body by using the ShouldBindJSON method and pass the schema
	if err := c.ShouldBindJSON(&input); err != nil {
		// If the data is invalid, it will return a 400 error to the client and tell them which fields are invalid.
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create a new book, save it to the database
	book := models.Book{Title: input.Title, Author: input.Author}
	models.DB.Create(&book)

	// Return the book
	c.JSON(http.StatusOK, gin.H{"data": book})
}

// GET /books/:id
// Find a specific book
func FindBook(c *gin.Context) {
	var book models.Book

	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// return the book.
	c.JSON(http.StatusOK, gin.H{"data": book})
}

// PATCH /books/:id
// Update a book by it's ID
func UpdateBook(c *gin.Context) {
	var book models.Book

	// Get the book if it exists
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate the input
	var input UpdatedBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errpr": err.Error()})
		return
	}

	models.DB.Model(&book).Update(input)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// DELETE /books/:id
// Delete a specific book based on it's ID
func DeleteBook(c *gin.Context) {
	// Get model if exist
	var book models.Book
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&book)

	c.JSON(http.StatusOK, gin.H{"deleted": true})
}
