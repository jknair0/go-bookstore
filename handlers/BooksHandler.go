package handlers

import (
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strings"
	"tech.jknair/bookstore/db"
	"tech.jknair/bookstore/mapper"
	"tech.jknair/bookstore/model"
	"tech.jknair/bookstore/response"
)

const RootRoute = "/"
const ItemRoute = "/{uuid:[a-z0-9-]+}/"

type BooksHandler struct {
	database   db.Database
	router     *mux.Router
	bookMapper *mapper.BookMapper
}

func CreateBookHandler(database db.Database, router *mux.Router) *BooksHandler {
	return &BooksHandler{
		database: database,
		router:   router,
	}
}

func (b *BooksHandler) Initialize() {
	b.router.HandleFunc(RootRoute, b.listBooks).Methods(http.MethodGet)
	b.router.HandleFunc(RootRoute, b.addBook).Methods(http.MethodPost)
	b.router.HandleFunc(ItemRoute, b.getBook).Methods(http.MethodGet)
}

func (b BooksHandler) getBook(writer http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uuid := vars["uuid"]
	writer.WriteHeader(http.StatusOK)
	if len(vars) == 0 || len(uuid) == 0 {
		writer.WriteHeader(http.StatusBadRequest)
		response.WriteErrorResponse(writer, response.ErrInvalidRoute)
		return
	}
	book := b.database.GetBook(uuid)
	if book == nil {
		response.WriteErrorResponse(writer, response.ErrItemNotFound)
		return
	}
	response.WriteSuccessResponse(writer, book)
}

func (b *BooksHandler) listBooks(writer http.ResponseWriter, _ *http.Request) {
	writer.WriteHeader(http.StatusOK)
	dbBookModelArray := b.database.GetBooks()
	bookArray := b.bookMapper.FromData(dbBookModelArray...)
	response.WriteSuccessResponse(writer, bookArray)
}

func (b *BooksHandler) addBook(w http.ResponseWriter, r *http.Request) {
	body := r.Body
	if body == nil {
		response.WriteErrorResponse(w, response.ErrInvalidRequestParams)
		return
	}
	requestBody, err := ioutil.ReadAll(body)
	if err != nil {
		response.WriteErrorResponse(w, response.ErrServerError)
		return
	}

	newBook, err := model.DecodeBook(requestBody)

	if err != nil {
		response.WriteErrorResponse(w, response.ErrInvalidRequestParams)
		return
	}

	if len(strings.TrimSpace(newBook.Name)) == 0 || len(strings.TrimSpace(newBook.Author)) == 0 {
		response.WriteErrorResponse(w, response.ErrInvalidRequestParams)
		return
	}

	w.WriteHeader(http.StatusOK)
	bookSchema := b.bookMapper.ToData(newBook)
	uidArray := b.database.SaveBooks(bookSchema)
	response.WriteSuccessResponse(w, uidArray)
}
