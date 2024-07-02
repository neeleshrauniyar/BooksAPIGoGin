package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "Welcome to the Books Store",
		})
	})

	router.GET("/book/get/books", GetBooks)
	router.GET("book/get/:id", GetABook)
	router.POST("book/create", CreateABook)
	router.PUT("book/update/:id", UpdateABook)
	router.DELETE("book/delete/:id", DeleteABook)

	router.Run()
}

type Book struct {
	ID          string `json: "id`
	Name        string `json: "name"`
	Price       int    `json: "price"`
	IsAvailable bool   `json: "isAvailable"`
}

var books = []Book{
	{ID: "1", Name: "Nepal Revisited", Price: 1500, IsAvailable: true},
	{ID: "2", Name: "China Unplugged", Price: 1000, IsAvailable: false},
	{ID: "3", Name: "Vietnam Exposed", Price: 500, IsAvailable: true},
}

func GetBooks(context *gin.Context) {
	context.JSON(200, books)
}

func GetABook(context *gin.Context) {
	id := context.Param("id")

	for _, book := range books {
		if book.ID == id {
			context.JSON(200, gin.H{
				"book": book,
			})
			return
		}
	}

	context.JSON(http.StatusNotFound, gin.H{
		"error": "The book does not exist",
	})
}

func CreateABook(context *gin.Context) {
	var newBook Book
	err := context.ShouldBindJSON(&newBook)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Book could not be created",
		})
		return
	}
	books = append(books, newBook)
	context.JSON(http.StatusCreated, gin.H{
		"message": "Book successfully created",
		"books":   books,
	})

}

func UpdateABook(context *gin.Context) {
	id := context.Param("id")

	for _, book := range books {
		if book.ID == id {
			book.Name = "Gulf Decoded"
			context.JSON(http.StatusOK, gin.H{
				"message":      "Book successfully updated",
				"updated-book": book,
			})
		}
	}
}

func DeleteABook(context *gin.Context) {
	id := context.Param("id")

	for index, book := range books {
		if book.ID == id {
			books = append(books[:index], books[index+1:]...)
			context.JSON(http.StatusOK, gin.H{
				"message": "Book sucessfully deleted",
				"books":   books,
			})
			return
		}
	}
	context.JSON(http.StatusBadRequest, gin.H{
		"message": "The Book could not be deleted",
	})
}
