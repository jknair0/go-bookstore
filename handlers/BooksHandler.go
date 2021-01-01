package handlers

import (
	"github.com/gorilla/mux"
	"github.com/jknair0/bookstore/db"
	"github.com/jknair0/bookstore/mapper"
	"github.com/jknair0/bookstore/model"
	"github.com/jknair0/bookstore/response"
	"io/ioutil"
	"net/http"
	"strings"
)

const ItemSlug = "uuid"

const RootRoute = "/"
const ItemRoute = "/{" + ItemSlug + ":[a-zA-Z0-9]{1,50}}"

const InvalidNameMessage = "Invalid Name"
const InvalidAuthorMessage = "Invalid Author"

type BooksHandler struct {
	router     *mux.Router
	database   db.Database
	bookMapper *mapper.BookMapper
}

func NewBooksHandler(router *mux.Router, database db.Database) *BooksHandler {
	return &BooksHandler{
		router:   router,
		database: database,
	}
}

func (b *BooksHandler) Initialize() {
	b.router.HandleFunc(RootRoute, b.listBooks).Methods(http.MethodGet)
	b.router.HandleFunc(RootRoute, b.addBook).Methods(http.MethodPost)
	b.router.HandleFunc(ItemRoute, b.getBook).Methods(http.MethodGet)
}

func (b BooksHandler) getBook(writer http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if len(vars) == 0 {
		writer.WriteHeader(http.StatusBadRequest)
		response.WriteErrorCodeResponse(writer, response.ErrInvalidRoute)
		return
	}
	uuid := vars[ItemSlug]
	if len(strings.TrimSpace(uuid)) == 0 {
		writer.WriteHeader(http.StatusBadRequest)
		response.WriteErrorCodeResponse(writer, response.ErrInvalidRoute)
		return
	}
	writer.WriteHeader(http.StatusOK)
	book := b.database.GetBook(uuid)
	if book == nil {
		response.WriteErrorCodeResponse(writer, response.ErrItemNotFound)
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
		response.WriteErrorCodeResponse(w, response.ErrInvalidRequestBody)
		return
	}
	requestBody, err := ioutil.ReadAll(body)
	if err != nil {
		response.WriteErrorCodeResponse(w, response.ErrServerError)
		return
	}
	newBook, err := model.DecodeBook(requestBody)
	if err != nil {
		response.WriteErrorCodeResponse(w, response.ErrInvalidRequestFormat)
		return
	}
	if len(strings.TrimSpace(newBook.Name)) == 0 {
		response.WriteErrorCodeCustomMessageResponse(w, response.ErrInvalidRequestFormat, InvalidNameMessage)
		return
	}
	if len(strings.TrimSpace(newBook.Author)) == 0 {
		response.WriteErrorCodeCustomMessageResponse(w, response.ErrInvalidRequestFormat, InvalidAuthorMessage)
		return
	}
	w.WriteHeader(http.StatusOK)
	bookSchema := b.bookMapper.ToData(newBook)
	uidArray := b.database.SaveBooks(bookSchema)
	response.WriteSuccessResponse(w, uidArray)
}
