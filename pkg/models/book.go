package models

import (
	"log"

	"gorm.io/gorm"
	// "gorm.io/gorm/logger"

	"github.com/simply-kk/GoBookStore/pkg/config"
)

var db *gorm.DB

// Book struct represents the database model
type Book struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

// Initialize the database connection
func init() {
	config.Connect()
	db = config.GetDB()

	// Ensure table creation with error handling
	if err := db.AutoMigrate(&Book{}); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
}

// CreateBook inserts a new book record into the database
func (b *Book) CreateBook() *Book {
	db.Create(b)
	return b
}

// GetAllBooks retrieves all books from the database
func GetAllBooks() []Book {
	var books []Book
	db.Find(&books)
	return books
}

// GetBookById retrieves a book by ID
func GetBookById(Id uint) (*Book, *gorm.DB) {
	var getBook Book
	dbResult := db.First(&getBook, Id)
	return &getBook, dbResult
}

// DeleteBook deletes a book by ID
func DeleteBook(ID uint) {
	var book Book
	db.Delete(&book, ID)
}
