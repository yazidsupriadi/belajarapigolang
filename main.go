package main

import (
	"belajarapigolang/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	//api route versioning
	v1 := router.Group("/v1")

	router.GET("/", handler.RouteHandler)
	v1.GET("/hello", handler.HelloHandler)
	v1.GET("/book/:id", handler.BooksHandler)
	v1.GET("/book", handler.BooksParamHandler)
	v1.POST("/bookpost", handler.PostBooksHandler)
	//atur route port
	router.Run(":1234")
}
