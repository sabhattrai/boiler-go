package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "science", Author: "sabin", Quantity: 5},
	{ID: "2", Title: "nepali", Author: "nabin", Quantity: 9},
	{ID: "3", Title: "social", Author: "rabin", Quantity: 2},
}

func geetbook(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}
func createbook(c *gin.Context) {
	var newBook book
	if err := c.BindJSON(&newBook); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}
func bookById(c *gin.Context) {
	id := c.Param("id")
	book, err := getBookById(id)
	if err != nil {
		return
	}
	c.IndentedJSON(http.StatusOK, book)

}
func getBookById(id string) (*book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}
	return nil, errors.New("book not found")
}
func main() {
	router := gin.Default()
	router.GET("/books", geetbook)
	router.GET("/books/:id", bookById)
	router.POST("/books", createbook)
	router.Run("localhost:8080")
}
