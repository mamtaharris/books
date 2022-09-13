package main

import (
	_ "github.com/mamtaharris/books/docs"
	"github.com/mamtaharris/books/routes"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title 			Bookstore API
// @version         1.0
// @description     A book management service API in Go using Gin framework.

// @contact.name   Mamta Harris
// @contact.url    https://google.com/
// @contact.email  mamta@gmail.com;harris@gmail.com

// @host      localhost:8080
// @BasePath  /v1
func main() {
	router := routes.SetupRouter()
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":8080")
}
