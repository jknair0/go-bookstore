package main

import (
	"fmt"
	"github.com/jknair0/bookstore/db"
	"github.com/jknair0/bookstore/routing"
	"log"
	"net/http"
)

var bookInMemoryDb = db.CreateInMemoryDb()

const PORT = "8000"

func main() {
	hostAddress := fmt.Sprintf(":%v", PORT)
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.Handle("/api/", routing.ApiRouter(bookInMemoryDb))
	err := http.ListenAndServe(hostAddress, nil)
	if err != nil {
		log.Fatalf("Error starting server %#v", err)
	}
}
