package handlers

import (
	"bytes"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"tech.jknair/bookstore/db/schema"
	"tech.jknair/bookstore/testutils"
	"testing"
)

var mockDatabase *testutils.MockDatabase
var handler *BooksHandler
var runTest func(func())

func TestMain(m *testing.M) {
	runTest = testutils.CreateForEach(setUp, tearDown)
	m.Run()
	runTest = nil
}

func setUp() {
	mockDatabase = testutils.CreateMockDatabase()
	handler = CreateBookHandler(mockDatabase, &mux.Router{})
}

func tearDown() {
	mockDatabase = nil
	handler = nil
}

func TestBooksHandler_AddBook(t *testing.T) {
	runTest(func() {
		requestBody := []byte(`{"name":"Deep Work","author":"Carl Jung"}`)
		r, _ := http.NewRequest(http.MethodPost, "/books", bytes.NewBuffer(requestBody))
		recorder := httptest.NewRecorder()

		expectedArg := []*schema.BookSchema{{
			Uuid:      "",
			Name:      "Deep Work",
			Author:    "Carl Jung",
			CreatedAt: 0,
		}}
		mockDatabase.On("SaveBooks", expectedArg).Return([]string{"123"})
		handler.addBook(recorder, r)
		responseBody := recorder.Body.Bytes()
		assert.Equal(t, http.StatusOK, recorder.Code, "invalid request code: %d", recorder.Code)
		assert.NotEqual(t, 0, len(responseBody), "empty response body")
		assert.Equal(t, string(responseBody), "[\"123\"]\n", "invalid response body: %s", string(responseBody))
	})
}

func TestBooksHandler_ListBooks(t *testing.T) {
	runTest(func() {
		r, _ := http.NewRequest(http.MethodGet, "/books", bytes.NewBuffer([]byte{}))
		recorder := httptest.NewRecorder()
		mockDatabase.On("GetBooks").Return([]*schema.BookSchema{})
		handler.listBooks(recorder, r)
		responseBody := recorder.Body.Bytes()
		assert.Equal(t, http.StatusOK, recorder.Code)
		assert.Equal(t, "[]", string(responseBody))
	})
}
