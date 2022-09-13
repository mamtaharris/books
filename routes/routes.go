package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mamtaharris/books/handlers"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/v1")
	v1.GET("/books", handlers.GetBooks)
	v1.GET("/books/:isbn", handlers.GetBookByISBN)
	v1.POST("/books", handlers.PostBook)

	return router
}
