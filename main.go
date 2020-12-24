package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"tech.jknair/bookstore/db"
	"tech.jknair/bookstore/handlers"
)

var (
	bookInMemoryDb = db.CreateInMemoryDb()
	bookDb         = bookInMemoryDb
)

const PORT = "8000"

func main() {
	hostAddress := fmt.Sprintf(":%v", PORT)
	router := getMuxRouter()
	http.Handle("/", router)
	err := http.ListenAndServe(hostAddress, nil)
	if err != nil {
		log.Fatalf("Error starting server %#v", err)
	}
}

func getMuxRouter() *mux.Router {
	muxRouter := mux.NewRouter()
	muxRouter.StrictSlash(true)
	muxRouter.Use(ContentTypeMiddleWare)
	muxRouter.HandleFunc("/", rootHandler)

	booksRouter := muxRouter.PathPrefix("/books").Subrouter()
	bookHandler := handlers.CreateBookHandler(bookDb, booksRouter)
	bookHandler.Initialize()

	return muxRouter
}

func rootHandler(writer http.ResponseWriter, _ *http.Request) {
	_, err := writer.Write([]byte("Welcome to BookStore"))
	if err != nil {
		log.Fatal("failed to handle root route")
	}
}
