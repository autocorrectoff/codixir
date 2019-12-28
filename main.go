package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	_ "strconv"

	"codixir/driver"
	"codixir/controllers"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/subosito/gotenv"
)

var db *sql.DB

func init() {
	gotenv.Load()
}

func main() {
	db = driver.Connect()
	router := mux.NewRouter()

	controller := controllers.BookController{}

	router.HandleFunc("/books", controller.GetBooks(db)).Methods("GET")
	router.HandleFunc("/books/{id}", controller.GetBook(db)).Methods("GET")
	router.HandleFunc("/books", controller.AddBook(db)).Methods("POST")
	router.HandleFunc("/books", controller.UpdateBook(db)).Methods("PUT")
	router.HandleFunc("/books/{id}", controller.RemoveBook(db)).Methods("DELETE")

	fmt.Println("Server is up")
	log.Fatal(http.ListenAndServe(":8000", router))
}
