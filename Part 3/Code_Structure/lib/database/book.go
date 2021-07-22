package database

import (
	"pebruwantoro/structuring/config"
	"pebruwantoro/structuring/models"
)

func GetBooks() (interface{}, error) {
	var books []models.Book
	if err := config.DB.Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

func GetOneBook(id int) (models.Book, error) {
	var book models.Book
	if err := config.DB.Find(&book, "id=?", id).Error; err != nil {
		return book, err
	}
	return book, nil
}

func CreateBook(book models.Book) (interface{}, error) {
	if err := config.DB.Save(&book).Error; err != nil {
		return nil, err
	}
	return book, nil
}

func DeleteBook(id int) (interface{}, error) {
	var book models.Book
	if err := config.DB.Where("id=?", id).First(&book).Error; err != nil {
		return nil, err
	}
	if err := config.DB.Delete(&book).Error; err != nil {
		return nil, err
	}
	return book, nil
}

func UpdateBook(book models.Book) (models.Book, error) {
	if err := config.DB.Save(&book).Error; err != nil {
		return book, err
	}
	return book, nil
}
