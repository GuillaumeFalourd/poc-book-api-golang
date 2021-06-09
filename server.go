package main

import (
	"log"
	"net/http"
	"time"

	"github.com/GuillaumeFalourd/poc-book-api-golang/datastore"
	"github.com/gorilla/mux"
)

var (
	books datastore.BookStore
)

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func init() {
	defer timeTrack(time.Now(), "file load")
	books = &datastore.Books{}
	books.Initialize()
}

// Add all routes here
func main() {
	log.Println("Start Book Rest API")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/books", createBook).Methods(http.MethodPost)
	router.HandleFunc("/books/authors/{author}", searchByAuthor).Methods(http.MethodGet)
	router.HandleFunc("/books/book-name/{bookName}", searchByBookName).Methods(http.MethodGet)
	router.HandleFunc("/books/isbn/{isbn}", searchByISBN).Methods(http.MethodGet)
	router.HandleFunc("/books/isbn/{isbn}", updateByISBN).Methods(http.MethodPut)
	router.HandleFunc("/books/isbn/{isbn}", deleteByISBN).Methods(http.MethodDelete)
	log.Fatalln(http.ListenAndServe(":8080", router))
}
