package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type book struct {
	ID     string  `json:"bookId"`
	Name   string  `json:"bookName"`
	Author string  `json:"authorName"`
	Price  float64 `json:"bookPrice"`
}

var books = []book{
	{
		ID:     "01",
		Name:   "TestBook",
		Author: "Khing",
		Price:  12.3,
	},
	{
		ID:     "02",
		Name:   "TestBook2",
		Author: "Khing2",
		Price:  14.6,
	},
	{
		ID:     "03",
		Name:   "TestBook3",
		Author: "Khing3",
		Price:  20.9,
	},
}

func getAllBooks(context *gin.Context) {
	context.JSON(http.StatusOK, books)
}

func getBookByName(context *gin.Context) {
	param := context.Param("name")
	var selectedBook []book

	for _, book := range books {
		if strings.EqualFold(book.Name, param) {
			selectedBook = append(selectedBook, book)
		}
	}

	context.JSON(http.StatusOK, selectedBook)
}

func addNewBook(context *gin.Context) {
	var newBook book
	//bind request body to editBook (pass as address of editBook in memory)
	if err := context.BindJSON(&newBook); err != nil {
		return //return on error
	}

	books = append(books, newBook)
	context.JSON(http.StatusOK, books)
}

func updateBookById(context *gin.Context) {
	var editBook book
	if err := context.BindJSON(&editBook); err != nil {
		return
	}

	param := context.Param("id")
	for index := range books {
		if books[index].ID == param {
			books[index].Name = editBook.Name
			books[index].Author = editBook.Author
			books[index].Price = editBook.Price

			context.JSON(http.StatusOK, books[index])
			return
		}
	}

	context.JSON(http.StatusNotFound, "data not found")
}

func deleteBookById(context *gin.Context) {
	param := context.Param("id")

	for index := range books {
		if books[index].ID == param {
			books = append(books[:index], books[index+1:]...)
			context.JSON(http.StatusOK, books)
			return
		}
	}
	context.JSON(http.StatusNotFound, "data not found")
}

func main() {
	router := gin.Default()

	router.GET("/books", getAllBooks)
	router.GET("/book/:name", getBookByName)
	router.POST("/books", addNewBook)
	router.PUT("/book/:id", updateBookById)
	router.DELETE("/book/:id", deleteBookById)

	router.Run("localhost:8000")
}
