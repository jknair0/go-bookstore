package routing

import (
	"github.com/gorilla/mux"
	"github.com/jknair0/bookstore/db"
	"github.com/jknair0/bookstore/handlers"
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
