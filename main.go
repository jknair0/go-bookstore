package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jknair0/bookstore/db"
	"github.com/jknair0/bookstore/handlers"
	"github.com/jknair0/bookstore/middleware"
	"log"
	"net/http"
)

var bookInMemoryDb = db.CreateInMemoryDb()

const PORT = "8000"

func main() {
	hostAddress := fmt.Sprintf(":%v", PORT)

	rootRouter := mux.NewRouter()
	rootRouter.Use(middleware.LoggingMiddleWare)
	rootRouter.StrictSlash(true)

	// static route
	staticFileServer := http.FileServer(http.Dir("./static"))
	rootRouter.Handle("/", staticFileServer)

	// '/api' route
	apiRouter := rootRouter.PathPrefix("/api").Subrouter()
	apiRouter.Use(middleware.ContentTypeMiddleWare)

	initHandlers(
		// '/books' route
		handlers.NewBooksHandler(apiRouter.PathPrefix("/books").Subrouter(), bookInMemoryDb),
	)

	http.Handle("/", rootRouter)
	err := http.ListenAndServe(hostAddress, nil)
	if err != nil {
		log.Fatalf("Error starting server %#v", err)
	}
}

func initHandlers(handlers ...handlers.AppHandler) {
	for _, handler := range handlers {
		handler.Initialize()
	}
}
