package routing

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"tech.jknair/bookstore/db"
	"tech.jknair/bookstore/handlers"
)

func GetRootRouter(bookDb db.Database) *mux.Router {
	muxRouter := mux.NewRouter()
	muxRouter.StrictSlash(true)
	muxRouter.HandleFunc("/", indexHandler)
	muxRouter.Use(LoggingMiddleWare)
	muxRouter.Use(ContentTypeMiddleWare)

	booksRouter := muxRouter.PathPrefix("/books").Subrouter()
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
