package main

import (
	"fmt"
	"log"
	"net/http"
	"tech.jknair/bookstore/db"
	"tech.jknair/bookstore/routing"
)

var bookInMemoryDb = db.CreateInMemoryDb()

const PORT = "8000"

func main() {
	hostAddress := fmt.Sprintf(":%v", PORT)
	http.Handle("/", routing.GetRootRouter(bookInMemoryDb))
	err := http.ListenAndServe(hostAddress, nil)
	if err != nil {
		log.Fatalf("Error starting server %#v", err)
	}
}
