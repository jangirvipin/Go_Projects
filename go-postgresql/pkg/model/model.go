package model

import (
	"github.com/jangirvipin/go-postgresql/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name      string `json:"name"`
	Number    string `json:"model"`
	Publisher string `json:"publisher"`
	AuthorID  uint   `json:"author_id"` // Foreign key
	Author    Author `json:"author" gorm:"foreignKey:AuthorID;references:ID"`
}

type Author struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	err := db.AutoMigrate(&Author{}, &Book{})
	if err != nil {
		panic(err.Error())
	}
}

func GetAllBooks() ([]Book, error) {
	var books []Book
	err := db.Find(&books)
	if err.Error != nil {
		return books, err.Error
	}
	return books, nil
}

func GetBookById(id int64) (*Book, error) {
	var book Book
	err := db.Where("id =?", id).First(&book)
	if err != nil {
		return nil, err.Error
	}
	return &book, nil
}

func DeleteBook(id int64) (*Book, error) {
	var book Book

	// First, get the book from DB
	if err := db.First(&book, id).Error; err != nil {
		return nil, err // not found or error
	}

	// Then delete the book
	if err := db.Delete(&book).Error; err != nil {
		return nil, err
	}

	return &book, nil
}

func (b *Book) CreateBook() (*Book, error) {
	// Create the author first if it exists and has no ID
	if b.Author.ID == 0 && b.Author.Name != "" {
		if err := db.Create(&b.Author).Error; err != nil {
			return nil, err
		}
		b.AuthorID = b.Author.ID
	}

	// Now create the book
	if err := db.Create(&b).Error; err != nil {
		return nil, err
	}
	return b, nil
}

func UpdateBook(book *Book) (*Book, error) {
	result := db.Save(book)
	if result.Error != nil {
		return nil, result.Error
	}
	return book, nil
}
