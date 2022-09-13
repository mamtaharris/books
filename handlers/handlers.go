package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mamtaharris/books/db"
	"github.com/mamtaharris/books/models"
)

// @summary      Get books array
// @description  Responds with the list of all books as JSON.
// @tags         books
// @produce      json
// @success      200  {array}  models.Book
// @router       /books [get]
func GetBooks(c *gin.Context) {
	c.JSON(http.StatusOK, db.Books)
}

// @summary      Store a new book
// @description  Takes a book JSON and store in DB. Return saved JSON.
// @tags         books
// @produce      json
// @param        book  body      models.Book  true  "Book JSON"
// @success      200	{object}  models.Book
// @router       /books [post]
func PostBook(c *gin.Context) {
	var newBook models.Book
	if err := c.BindJSON(&newBook); err != nil {
		return
	}
	db.Books = append(db.Books, newBook)
	c.JSON(http.StatusCreated, newBook)
}

// @summary      Get single book by isbn
// @description  Returns the book whose ISBN value matches the isbn.
// @tags         books
// produce      json
// @param        isbn  path      string  true  "search book by isbn"
// @success      200  {object}  models.Book
// @router       /books/{isbn} [get]
func GetBookByISBN(c *gin.Context) {
	isbn := c.Param("isbn")
	for _, a := range db.Books {
		if a.ISBN == isbn {
			c.JSON(http.StatusOK, a)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "book not found"})
}
