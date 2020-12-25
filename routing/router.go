package routing

import (
	"github.com/gorilla/mux"
	"github.com/jknair0/bookstore/db"
	"github.com/jknair0/bookstore/handlers"
	"log"
	"net/http"
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
