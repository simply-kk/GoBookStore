package models

import (
	"github.com/jinzhu/gorm"
	"github.com/simply-kk/GoBookStore/pkg/config"
)

var db *gorm.DB

// Book struct represents the database model
type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

// Initialize the database connection
func init() {
	config.Connect()
	db = config.GetDB()

	// Ensure table creation with error handling
	if err := db.AutoMigrate(&Book{}).Error; err != nil {
		panic("Failed to migrate database: " + err.Error())
	}
}

// CreateBook inserts a new book record into the database
func (b *Book) CreateBook() *Book {
	db.Create(b)
	return b
}

// GetAllBooks retrieves all books from the database
func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

// GetBookById retrieves a book by ID
func GetBookById(Id int64) (*Book, *gorm.DB) {
	var GetBook Book
	dbResult := db.Where("ID=?", Id).Find(&GetBook)
	return &GetBook, dbResult
}

// DeleteBook deletes a book by ID
func DeleteBook(ID int64) Book {
	var book Book
	db.Where("ID=?", ID).Delete(&book)
	return book
}
