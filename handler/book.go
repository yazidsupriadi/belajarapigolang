package handler

import (
	"belajarapigolang/book"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type bookHandler struct {
	bookService book.Service
}

func NewBookHandler(bookService book.Service) *bookHandler {
	return &bookHandler{bookService}
}

//handler function
/*
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
*/

func (h *bookHandler) GetBooksHandler(c *gin.Context) {
	books, err := h.bookService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}
	var booksResponse []book.BookResponse

	for _, b := range books {
		bookResponse := book.BookResponse{
			ID:          b.ID,
			Title:       b.Title,
			Price:       b.Price,
			Description: b.Description,
			Rating:      b.Rating,
		}
		booksResponse = append(booksResponse, bookResponse)
	}
	c.JSON(http.StatusOK, gin.H{
		"data": booksResponse,
	})

}
func (h *bookHandler) PostBooksHandler(c *gin.Context) {

	var bookRequest book.BookRequest

	err := c.ShouldBindJSON(&bookRequest)

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

	book, err := h.bookService.Create(bookRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"data": book,
	})
}

func (h *bookHandler) GetBookHandler(c *gin.Context) {

	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	b, err := h.bookService.FindById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}
	bookResponse := book.BookResponse{
		ID:          b.ID,
		Title:       b.Title,
		Price:       b.Price,
		Description: b.Description,
		Rating:      b.Rating,
	}

	c.JSON(http.StatusOK, gin.H{
		"data": bookResponse,
	})
}

func (h *bookHandler) UpdateBookHandler(c *gin.Context) {

	var bookRequest book.BookRequest

	err := c.ShouldBindJSON(&bookRequest)

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

	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	book, err := h.bookService.Update(id, bookRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"data": book,
	})
}

func (h *bookHandler) DeleteBookHandler(c *gin.Context) {

	var bookRequest book.BookRequest

	err := c.ShouldBindJSON(&bookRequest)

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

	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	book, err := h.bookService.Delete(int(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"data": book,
	})
}
