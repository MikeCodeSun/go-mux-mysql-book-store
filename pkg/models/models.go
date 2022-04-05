package models

import (
	"example/03-mux-gorm-mysql-book/pkg/config"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name string `json:"name"`
	Author string `json:"author"`
	Publication string `json:"publication"`
}

var db *gorm.DB

func init() {
	config.ConnectDB()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func GetAllBooks() []Book {
var Books []Book
	db.Find(&Books)
	return Books
} 

func GetOneBook(id int64) (*Book, *gorm.DB) {
	var Book Book
	db := db.Where("ID = ?", id).Find(&Book)
	return &Book ,db
}

func (B *Book) CreateBook() *Book {
	db.Create(&B)
	return B
}

func DeleteBook(id int64) Book {
	var book Book
	db.Where("ID = ?", id).Delete(&book)
	return book
}