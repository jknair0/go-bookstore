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

func TestBooksHandler_AddBook(t *testing.T) {
	requestBody := []byte(`{"name":"Deep Work","author":"Carl Jung"}`)
	r, _ := http.NewRequest(http.MethodGet, "/books", bytes.NewBuffer(requestBody))
	recorder := httptest.NewRecorder()

	mockDatabase := testutils.CreateMockDatabase()
	expectedArg := []*schema.BookSchema{{
		Uuid:      "",
		Name:      "Deep Work",
		Author:    "Carl Jung",
		CreatedAt: 0,
	}}
	mockDatabase.On("SaveBooks", expectedArg).Return([]string{"123"})

	handler := CreateBookHandler(mockDatabase, &mux.Router{})
	handler.addBook(recorder, r)

	responseBody := recorder.Body.Bytes()

	assert.Equal(t, http.StatusOK, recorder.Code, "invalid request code: %d", recorder.Code)
	assert.NotEqual(t, 0, len(responseBody), "empty response body")
	assert.Equal(t, string(responseBody), "[\"123\"]\n", "invalid response body: %s", string(responseBody))
}
