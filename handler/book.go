package handler

import (
	"belajarapigolang/book"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

//handler function
func RouteHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "yazid",
		"bio":  "Back end engineer",
	})
}

func HelloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"title":   "Hello",
		"content": "Hi we can say hello",
	})
}

func BooksHandler(c *gin.Context) {

	id := c.Param("id")

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func BooksParamHandler(c *gin.Context) {

	title := c.Query("title")

	c.JSON(http.StatusOK, gin.H{
		"title": title,
	})
}

func PostBooksHandler(c *gin.Context) {

	var bookInput book.BookInput

	err := c.ShouldBindJSON(&bookInput)

	if err != nil {
		//error stop server
		//log.Fatal(err)

		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("error on field %s,condition %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)

		}
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errorMessages,
		})
		fmt.Println(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"title":    bookInput.Title,
		"price":    bookInput.Price,
		"subtitle": bookInput.SubTitle,
	})
}
