package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"tech.jknair/bookstore/db"
	"tech.jknair/bookstore/mapper"
	"tech.jknair/bookstore/model"
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
	if len(vars) == 0 || len(uuid) == 0 {
		writer.WriteHeader(http.StatusBadRequest)
		_, _ = writer.Write(model.CreateErrorResponseHolder(model.INVALID_PATH_ERROR_MESSAGE).EncodeJson())
		return
	}
	writer.WriteHeader(http.StatusOK)
	book := b.database.GetBook(uuid)
	if book == nil {
		_, _ = writer.Write(model.CreateErrorResponseHolder(model.ITEM_NOT_FOUND_ERROR_MESSAGE).EncodeJson())
		return
	}
	_, _ = writer.Write(model.CreateSuccessResponseHolder(book).EncodeJson())
}

func (b *BooksHandler) listBooks(writer http.ResponseWriter, _ *http.Request) {
	writer.WriteHeader(http.StatusOK)
	books := b.database.GetBooks()
	schemasArray := b.bookMapper.FromData(books...)
	_, _ = writer.Write(model.CreateSuccessResponseHolder(schemasArray).EncodeJson())
}

func (b *BooksHandler) addBook(w http.ResponseWriter, r *http.Request) {
	requestBodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "failed to read requestBodyBytes", http.StatusInternalServerError)
		return
	}

	if len(requestBodyBytes) == 0 {
		http.Error(w, "No request params received", http.StatusBadRequest)
		return
	}

	newBook, err := model.DecodeBook(requestBodyBytes)

	if err != nil {
		http.Error(w, "Invalid request params", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	bookSchema := b.bookMapper.ToData(newBook)
	uidArray := b.database.SaveBooks(bookSchema)
	_ = json.NewEncoder(w).Encode(uidArray)
}
