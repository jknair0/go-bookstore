package routing

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"tech.jknair/bookstore/db"
	"tech.jknair/bookstore/handlers"
)

func ApiRouter(bookDb db.Database) *mux.Router {
	muxRouter := mux.NewRouter()
	muxRouter.StrictSlash(true)
	muxRouter.Use(LoggingMiddleWare)

	booksRouter := muxRouter.PathPrefix("/api/books").Subrouter()
	booksRouter.Use(ContentTypeMiddleWare)
	bookHandler := handlers.CreateBookHandler(bookDb, booksRouter)
	bookHandler.Initialize()

	return muxRouter
}

func indexHandler(writer http.ResponseWriter, _ *http.Request) {
	_, err := writer.Write([]byte("Welcome to BookStore"))
	if err != nil {
		log.Fatal("failed to handle root route")
	}
}
