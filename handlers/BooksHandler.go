package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"tech.jknair/bookstore/db"
	"tech.jknair/bookstore/mapper"
	"tech.jknair/bookstore/model"
)

type BooksHandler struct {
	database   db.Database
	router     *mux.Router
	bookMapper mapper.BookMapper
}

func CreateBookHandler(database db.Database, router *mux.Router) BooksHandler {
	return BooksHandler{
		database: database,
		router:   router,
	}
}

func (b *BooksHandler) Initialize() {
	b.router.HandleFunc("/", b.listBooks).Methods(http.MethodGet)
	b.router.HandleFunc("/", b.addBook).Methods(http.MethodPost)
}

func (b *BooksHandler) listBooks(writer http.ResponseWriter, _ *http.Request) {
	writer.WriteHeader(http.StatusOK)
	books := b.database.GetBooks()
	schemasArray := b.bookMapper.FromData(books...)
	jsonBooks, err := schemasArray.EncodeBooks()
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, _ = writer.Write(jsonBooks)
}

func (b *BooksHandler) addBook(w http.ResponseWriter, r *http.Request) {
	requestBodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "failed to read requestBodyBytes", http.StatusInternalServerError)
		return
	}

	log.Printf("received: %s", requestBodyBytes)
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
