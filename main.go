package main

import (
	"strconv"
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

// Book model
type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   string `json:"year"`
}

var books []Book

func main() {
	router := mux.NewRouter()
	router.Headers("Content-Type", "application/json")

	books = append(books, Book{ID: 1, Title: "Clean Code: A Handbook of Agile Software Craftsmanship", Author: "Robert C. Martin", Year: "2008"},
		Book{ID: 2, Title: "Spring Microservices in Action", Author: "John Carnell", Year: "2017"},
		Book{ID: 3, Title: "Go Programming Blueprints", Author: "Mat Ryer", Year: "2016"})

	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", addBook).Methods("POST")
	router.HandleFunc("/books", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", removeBook).Methods("DELETE")

	fmt.Println("Server is up")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func getBooks(writer http.ResponseWriter, request *http.Request) {
	json.NewEncoder(writer).Encode(books)
}
func getBook(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, err := strconv.Atoi(params["id"])
	fmt.Println(err)
	if err == nil {
		for _, book := range books{
			if book.ID == id {
				json.NewEncoder(writer).Encode(&book)
			}
		}
	} else{
		log.Fatal(err)
	}
	
}
func addBook(writer http.ResponseWriter, request *http.Request) {
	var book Book
	json.NewDecoder(request.Body).Decode(&book)
	books = append(books, book)
	json.NewEncoder(writer).Encode(books)
}
func updateBook(writer http.ResponseWriter, request *http.Request) {
	var book Book
	json.NewDecoder(request.Body).Decode(&book)
	for i, item := range books{
		if book.ID == item.ID {
			books[i] = book
		}
	}
	json.NewEncoder(writer).Encode(books)
}
func removeBook(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, err := strconv.Atoi(params["id"])
	if err == nil {
		for i, book := range books{
			if book.ID == id {
				books = append(books[:i], books[i + 1:]...)
			}
		}
	} else{
		log.Fatal(err)
	}
	json.NewEncoder(writer).Encode(books)
}
