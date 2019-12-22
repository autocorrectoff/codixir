package main

import(
	"fmt"
	"log"
	"net/http"
	//"encoding/json"
	"github.com/gorilla/mux"
)

type Book struct{
	ID int `json:id`
	Title string `json:id`
	Author string `json:id`
	Year string `json:id`
}

var books []Book

func main(){
	router := mux.NewRouter()

	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", addBook).Methods("POST")
	router.HandleFunc("/books", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", removeBook).Methods("DELETE")

	fmt.Println("Server is up")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func getBooks(writer http.ResponseWriter, request *http.Request){
	log.Println("Getting all books")
}
func getBook(writer http.ResponseWriter, request *http.Request){
	log.Println("Getting one book")
}
func addBook(writer http.ResponseWriter, request *http.Request){
	log.Println("Getting one book")
}
func updateBook(writer http.ResponseWriter, request *http.Request){
	log.Println("Getting one book")
}
func removeBook(writer http.ResponseWriter, request *http.Request){
	log.Println("Getting one book")
}