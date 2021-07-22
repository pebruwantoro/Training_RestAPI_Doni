package controllers

import (
	"net/http"
	"pebruwantoro/middleware/config"
	"pebruwantoro/middleware/lib/database"
	"pebruwantoro/middleware/models"
	"strconv"

	"github.com/labstack/echo"
)

// get all books
func GetBooksControllers(c echo.Context) error {
	books, err := database.GetBooks()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "can not fetch data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"books":  books,
	})
}

// get book by id
func GetBookControllers(c echo.Context) error {
	id, e := strconv.Atoi(c.Param("id"))
	if e != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	var count int64
	config.DB.Model(&models.Book{}).Where("id=?", id).Count(&count)
	if count == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "not found",
		})
	}
	getBook, err := database.GetOneBook(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "can not fetch data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success get book by id",
		"books":  getBook,
	})
}

// create new book
func CreateBookControllers(c echo.Context) error {
	var book models.Book
	c.Bind(&book)
	newBook, err := database.CreateBook(book)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "can not fetch data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success create new book",
		"users":  newBook,
	})
}

func DeleteBookControllers(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	delete_book, err := database.DeleteBook(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "can not fetch data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success delete book",
		"book":   delete_book,
	})
}

func UpdateBookControllers(c echo.Context) error {
	var book models.Book
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	get_book, _ := database.GetOneBook(id)
	book = get_book
	c.Bind(&book)
	update_book, err := database.UpdateBook(book)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "can not fetch data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success update book",
		"users":  update_book,
	})
}
