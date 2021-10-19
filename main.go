package main

import (
	"belajarapigolang/book"
	"belajarapigolang/handler"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	dsn := "root:@tcp(127.0.0.1:3306)/latihanapigolang?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("DB Connection error")
	}
	fmt.Println("Database Connected!")
	db.AutoMigrate(&book.Book{})

	//find all with repository
	bookRepository := book.NewRepository(db)

	bookService := book.NewService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)

	/*
		books, err := bookRepository.FindAll()
		for _, book := range books {
			fmt.Println("Title", book.Title)
		}
	*/

	//crud

	//create or store data
	/*
		book := book.Book{}
		book.Title = "Atomic Habits"
		book.Price = 100000
		book.Rating = 5
		book.Description = "Buku self development"

		err = db.Create(&book).Error
		if err != nil {
			fmt.Println("===========")

			fmt.Println("There is an error occured!")
			fmt.Println("===========")
		}
	*/

	/*
		//single object

		var book book.Book

		//for list
		//var books []book.Book

		err = db.Debug().Where("id = ?", 1).First(&book).Error
		if err != nil {
			fmt.Println("===========")

			fmt.Println("There is an error occured!")
			fmt.Println("===========")
		}
		/*

		/*
			fmt.Println("===========")

			fmt.Println("Book Title : ", book.Title)
			fmt.Println("Book Price %w : ", book)
			fmt.Println("===========")
	*/

	/*
		looping like foreach of list

		for _, b := range books {
			fmt.Println("Book Title : ", b.Title)
			fmt.Println("Book Price %w : ", books)
			fmt.Println("===========")

		}
	*/

	/*

		//updated data

		book.Title = "Atomic Habits Revised Edition"
		err = db.Save(&book).Error
		if err != nil {
			fmt.Println("===========")

			fmt.Println("There is an error occured when updating record!")
			fmt.Println("===========")
		}
	*/

	/*
		err = db.Delete(&book).Error
		if err != nil {
			fmt.Println("===========")

			fmt.Println("There is an error occured when updating record!")
			fmt.Println("===========")
		}
	*/
	router := gin.Default()
	//api route versioning
	v1 := router.Group("/v1")

	/*
		router.GET("/", handler.RouteHandler)
		v1.GET("/hello", handler.HelloHandler)
		v1.GET("/book/:id", handler.BooksHandler)
		v1.GET("/book", handler.BooksParamHandler)
	*/
	v1.GET("/books", bookHandler.GetBooksHandler)
	v1.GET("/book/:id", bookHandler.GetBookHandler)
	v1.PUT("/book/update/:id", bookHandler.UpdateBookHandler)
	v1.DELETE("/book/delete/:id", bookHandler.DeleteBookHandler)
	v1.POST("/bookpost", bookHandler.PostBooksHandler)
	//atur route port
	router.Run(":1406")
}
