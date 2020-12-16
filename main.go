package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const PORT = "8000"

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
