package usecases

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"tech.jknair/bookstore/db"
	"tech.jknair/bookstore/usecases/model"
)

var (
	bookInMemoryDb = db.Create()
	bookDb    = &bookInMemoryDb
)

func AddBook(w http.ResponseWriter, r *http.Request) {
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "failed to read bytes", http.StatusInternalServerError)
		return
	}
	newBook := model.EmptyBook()
	err = json.Unmarshal(bytes, &newBook)
	if err != nil {
		http.Error(w, "failed to unmarshall", http.StatusInternalServerError)
		return
	}
	newBook := bookDb.SaveBooks([]db.BookSchema{ })
	_ = json.NewEncoder(w).Encode(newBook)
}

func ListBooks(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(books)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}
