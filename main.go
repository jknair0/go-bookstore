package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const PORT = "8000"

type Book struct {
	Name      string `json:"name"`
	Author    string `json:"author"`
	CreatedAt int64  `json:"created_at"`
}

func createBook(name string, author string) Book {
	return Book{
		Name:      name,
		Author:    author,
		CreatedAt: time.Now().Unix(),
	}
}

func emptyBook() Book {
	return createBook("", "")
}

var books = []Book{
	createBook("Deep work", "Carl Jung"),
	createBook("The Art of Computer Programming", "Donald Knuth"),
}

func addBook(w http.ResponseWriter, r *http.Request) {
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "failed to read bytes", http.StatusInternalServerError)
		return
	}
	newBook := emptyBook()
	err = json.Unmarshal(bytes, &newBook)
	if err != nil {
		http.Error(w, "failed to unmarshall", http.StatusInternalServerError)
		return
	}
	newSet := append(books, newBook)
	books = newSet
	json.NewEncoder(w).Encode(books)
}

func listBooks(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(books)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func contentTypeMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(writer, request)
	})
}

func main() {
	hostAddress := fmt.Sprintf(":%v", PORT)

	router := mux.NewRouter()

	router.HandleFunc("/", rootHandler)
	router.HandleFunc("/books", listBooks).Methods(http.MethodGet)
	router.HandleFunc("/books", addBook).Methods(http.MethodPost)

	router.Use(contentTypeMiddleWare)

	http.Handle("/", router)

	err := http.ListenAndServe(hostAddress, nil)
	if err != nil {
		log.Fatalf("Error starting server %#v", err)
	}
}

func rootHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Hello world"))
}