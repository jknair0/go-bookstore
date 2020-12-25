package main

import (
	"fmt"
	muxHandlers "github.com/gorilla/handlers"
	"github.com/jknair0/bookstore/db"
	"github.com/jknair0/bookstore/routing"
	"log"
	"net/http"
	"os"
)

var bookInMemoryDb = db.CreateInMemoryDb()

const PORT = "8000"

func main() {
	hostAddress := fmt.Sprintf(":%v", PORT)

	staticFileServer := http.FileServer(http.Dir("./static"))
	staticFileServer = muxHandlers.LoggingHandler(os.Stdout, staticFileServer)
	http.Handle("/", staticFileServer)

	http.Handle("/api/", routing.ApiRouter(bookInMemoryDb))

	err := http.ListenAndServe(hostAddress, nil)
	if err != nil {
		log.Fatalf("Error starting server %#v", err)
	}
}
