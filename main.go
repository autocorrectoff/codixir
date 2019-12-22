package main

import(
	"fmt"
	"log"
	"net/http"
	"encoding/json"
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

}