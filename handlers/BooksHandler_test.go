package handlers

import (
	"bytes"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"tech.jknair/bookstore/db/schema"
	"tech.jknair/bookstore/model"
	"tech.jknair/bookstore/testutils"
	"testing"
)

var mockDatabase *testutils.MockDatabase
var subject *BooksHandler
var runTest func(func())

func TestMain(m *testing.M) {
	runTest = testutils.CreateForEach(setUp, tearDown)
	m.Run()
	runTest = nil
}

func setUp() {
	mockDatabase = testutils.CreateMockDatabase()
	subject = CreateBookHandler(mockDatabase, &mux.Router{})
	subject.Initialize()
}

func tearDown() {
	mockDatabase = nil
	subject = nil
}

func TestBooksHandler_AddBook(t *testing.T) {
	runTest(func() {
		requestBody := []byte(`{"name":"Deep Work","author":"Carl Jung"}`)
		r, _ := http.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		recorder := httptest.NewRecorder()

		expectedArg := []*schema.BookSchema{{
			Uuid:      "",
			Name:      "Deep Work",
			Author:    "Carl Jung",
			CreatedAt: 0,
		}}
		mockDatabase.On("SaveBooks", expectedArg).Return([]string{"123"})
		subject.addBook(recorder, r)
		responseBody := recorder.Body.Bytes()
		assert.Equal(t, http.StatusOK, recorder.Code, "invalid request code: %d", recorder.Code)
		assert.NotEqual(t, 0, len(responseBody), "empty response body")
		assert.Equal(t, string(responseBody), "[\"123\"]\n", "invalid response body: %s", string(responseBody))
	})
}

func TestBooksHandler_GetBooks(t *testing.T) {
	runTest(func() {
		r, _ := http.NewRequest(http.MethodGet, "/", bytes.NewBuffer([]byte{}))
		recorder := httptest.NewRecorder()
		mockDatabase.On("GetBooks").Return([]*schema.BookSchema{})
		subject.listBooks(recorder, r)
		responseBody := recorder.Body.Bytes()
		assert.Equal(t, http.StatusOK, recorder.Code)
		assert.JSONEq(t, `{"data":[],"error":null}`, string(responseBody))
	})
}

func TestBooksHandler_GetBook(t *testing.T) {
	t.Run("invalid url", func(t *testing.T) {
		runTest(func() {
			validUuid := "invalid_path_url"
			url := fmt.Sprintf("/%s/", validUuid)
			req, err := http.NewRequest(http.MethodGet, url, nil)
			if err != nil {
				t.Fatal(err)
			}
			recorder := httptest.NewRecorder()

			subject.router.ServeHTTP(recorder, req)

			assert.Equal(t, recorder.Body.String(), "404 page not found\n")
		})
	})

	t.Run("valid url with valid uuid", func(t *testing.T) {
		runTest(func() {
			validUuid := "a-valid-uuid"
			url := fmt.Sprintf("/%s/", validUuid)
			req, err := http.NewRequest(http.MethodGet, url, nil)
			if err != nil {
				t.Fatal(err)
			}
			recorder := httptest.NewRecorder()
			sampleBookSchema := sampleBookSchema()
			mockDatabase.On("GetBook", validUuid).Return(sampleBookSchema)

			subject.router.ServeHTTP(recorder, req)

			expectedCalls := mockDatabase.ExpectedCalls
			assert.Len(t, expectedCalls, 1)
			getBookCall := expectedCalls[0]
			assert.Equal(t, validUuid, getBookCall.Arguments.String(0))
			assert.Equal(t, recorder.Code, http.StatusOK)
			assert.JSONEq(t, `{"data":{"Uuid":"","Name":"Deep Work","Author":"Carl Jung","CreatedAt":0},"error":null}`, recorder.Body.String())
		})
	})

	t.Run("valid uid with invalid uid", func(t *testing.T) {
		runTest(func() {
			invalidUid := "a-valid-uuid"
			url := fmt.Sprintf("/%s/", invalidUid)
			req, err := http.NewRequest(http.MethodGet, url, nil)
			if err != nil {
				t.Fatal(err)
			}
			recorder := httptest.NewRecorder()
			mockDatabase.On("GetBook", invalidUid).Return(nil)

			subject.router.ServeHTTP(recorder, req)
			expectedResponse := fmt.Sprintf(`{"data":null,"error":"%s"}`, model.ITEM_NOT_FOUND_ERROR_MESSAGE)
			assert.JSONEq(t, expectedResponse, recorder.Body.String())
		})
	})
}

func sampleBookSchema() *schema.BookSchema {
	return &schema.BookSchema{
		Uuid:      "",
		Name:      "Deep Work",
		Author:    "Carl Jung",
		CreatedAt: 0,
	}
}
