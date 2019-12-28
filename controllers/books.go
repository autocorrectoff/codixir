package controllers

import (
	"codixir/models"
	"strconv"
	"codixir/repository"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)
 // BookController -> name is s self explanatory
type BookController struct{}

var books []models.Book

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// GetBooks -> get all books
func (c BookController) GetBooks(db *sql.DB) http.HandlerFunc{
	return func(writer http.ResponseWriter, request *http.Request) {
		var book models.Book
		books = []models.Book{}
		repo := repository.BookRepository{}
		books = repo.GetBooks(db, book, books)
		json.NewEncoder(writer).Encode(books)
	}
}

// GetBook -> get book by id
func (c BookController) GetBook(db *sql.DB) http.HandlerFunc{
	return func (writer http.ResponseWriter, request *http.Request) {
		var book models.Book
		params := mux.Vars(request)
		repo := repository.BookRepository{}
		id, err := strconv.Atoi(params["id"])
		logFatal(err)
		book = repo.GetBook(db, book, id)
		json.NewEncoder(writer).Encode(book)
	}
}

// AddBook -> add one book
func (c BookController) AddBook(db *sql.DB) http.HandlerFunc{
	return func (writer http.ResponseWriter, request *http.Request) {
		var book models.Book
		json.NewDecoder(request.Body).Decode(&book)
		repo := repository.BookRepository{}
		bookID := repo.AddBook(db, book)
		json.NewEncoder(writer).Encode(bookID)
	}
}

// UpdateBook -> update book by id
func (c BookController) UpdateBook(db *sql.DB) http.HandlerFunc{
	return func(writer http.ResponseWriter, request *http.Request) {
		var book models.Book
		json.NewDecoder(request.Body).Decode(&book)
		repo := repository.BookRepository{}
		updated := repo.UpdateBook(db, book)
		json.NewEncoder(writer).Encode(updated)
	}
}

// RemoveBook -> delete book by id
func (c BookController) RemoveBook(db *sql.DB) http.HandlerFunc{
return func (writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	repo := repository.BookRepository{}
	id, err := strconv.Atoi(params["id"])
	logFatal(err)
	rowsDeleted := repo.RemoveBook(db, id)
	json.NewEncoder(writer).Encode(rowsDeleted)
}
}