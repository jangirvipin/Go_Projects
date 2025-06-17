package model

import (
	"github.com/jangirvipin/go-mysql/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name      string `json:"name"`
	Author    string `json:"author"`
	Publisher string `json:"publisher"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	err := db.AutoMigrate(&Book{})
	if err != nil {
		panic(err)
	}
}

func (b *Book) CreateBook() *Book {
	db.Create(&b)
	return b
}

func GetBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(ID int64) (*Book, *gorm.DB) {
	var getBook Book
	db.Where("id=?", ID).Find(&getBook)
	return &getBook, db
}

func DeleteBook(ID int64) *Book {
	var book Book
	db.Where("id=?", ID).Delete(&book)
	return &book
}

func UpdateBook(b *Book) *Book {
	db.Save(b)
	return b
}
