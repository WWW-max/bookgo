package controllers

import (
	"bookgo/database"
	"bookgo/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context) {
	rows, err := database.DB.Query("SELECT id, title, author, isbn FROM books")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch books"})
		return
	}

	defer rows.Close()

	var books []models.Book
	for rows.Next() {
		var book models.Book
		rows.Scan(&book.ID, &book.Title, &book.Author, &book.ISBN)
		books = append(books, book)
	}

	c.JSON(http.StatusOK, books)
}

func AddBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	_, err := database.DB.Exec("INSERT INTO books (title, author, isbn) VALUES (?, ?, ?)", book.Title, book.Author, book.ISBN)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add book"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Book added successfully"})
}
func UpdateBook(c *gin.Context) {
	id := c.Param("id")
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	_, err := database.DB.Exec("UPDATE books SET title = ?, author = ?, isbn = ? WHERE id = ?", book.Title, book.Author, book.ISBN, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update book"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book updated successfully"})
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")

	_, err := database.DB.Exec("DELETE FROM books WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete book"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}


	