package handlers

import (
	"bytes"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	beforeEach "github.com/jknair0/beforeeach"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"tech.jknair/bookstore/contants"
	"tech.jknair/bookstore/db/schema"
	"tech.jknair/bookstore/testutils"
	"testing"
)

var mockDatabase *testutils.MockDatabase
var subject *BooksHandler
var it = beforeEach.Create(setUp, tearDown)

func setUp() {
	mockDatabase = testutils.CreateMockDatabase()
	muxRouter := mux.NewRouter()
	subject = CreateBookHandler(mockDatabase, muxRouter)
	subject.Initialize()
}

func tearDown() {
	mockDatabase = nil
	subject = nil
}

func TestBooksHandler_AddBook(t *testing.T) {
	t.Run("valid request body", func(t *testing.T) {
		it(func() {
			requestBody := bytes.NewBufferString(`{"name":"Deep Work","author":"Carl Jung"}`)
			r, _ := http.NewRequest(http.MethodPost, RootRoute, requestBody)
			recorder := httptest.NewRecorder()
			fakeUid := uuid.New().String()
			expectedArg := []*schema.BookSchema{sampleBookSchema()}
			mockDatabase.On("SaveBooks", expectedArg).Return([]string{fakeUid})

			subject.router.ServeHTTP(recorder, r)

			responseBody := recorder.Body.String()
			assert.Equal(t, http.StatusOK, recorder.Code)
			expectedResponse := fmt.Sprintf(`["%s"]`, fakeUid)
			assert.JSONEq(t, responseBody, expectedResponse)
		})
	})
	t.Run("empty request body", func(t *testing.T) {
		it(func() {
			r, _ := http.NewRequest(http.MethodPost, RootRoute, nil)
			recorder := httptest.NewRecorder()

			subject.router.ServeHTTP(recorder, r)

			responseBody := recorder.Body.String()
			expectedResponse := fmt.Sprintf(`{"data":null,"error":"%s"}`, contants.ErrInvalidRequestParams)
			assert.JSONEq(t, expectedResponse, responseBody)
		})
	})
	t.Run("invalid request body", func(t *testing.T) {
		it(func() {
			requestBody := bytes.NewBufferString(`{"not_name":"sample","not_author":"sample-author"}`)
			r, _ := http.NewRequest(http.MethodPost, RootRoute, requestBody)
			recorder := httptest.NewRecorder()

			subject.router.ServeHTTP(recorder, r)

			responseBody := recorder.Body.String()
			expectedResponse := fmt.Sprintf(`{"data":null,"error":"%s"}`, contants.ErrInvalidRequestParams)
			assert.JSONEq(t, expectedResponse, responseBody)
		})
	})
}

func TestBooksHandler_GetBooks(t *testing.T) {
	t.Run("empty books", func(t *testing.T) {
		it(func() {
			r, _ := http.NewRequest(http.MethodGet, RootRoute, nil)
			recorder := httptest.NewRecorder()
			mockDatabase.On("GetBooks").Return([]*schema.BookSchema{})

			subject.router.ServeHTTP(recorder, r)

			responseBody := recorder.Body.String()
			assert.Equal(t, http.StatusOK, recorder.Code)
			assert.JSONEq(t, `{"data":[],"error":null}`, responseBody)
		})
	})
	t.Run("non empty books", func(t *testing.T) {
		it(func() {
			r, _ := http.NewRequest(http.MethodGet, RootRoute, nil)
			recorder := httptest.NewRecorder()
			mockDatabase.On("GetBooks").Return([]*schema.BookSchema{sampleBookSchema()})

			subject.router.ServeHTTP(recorder, r)

			responseBody := recorder.Body.String()
			assert.Equal(t, http.StatusOK, recorder.Code)
			assert.JSONEq(t, `{"data":[{"uuid":"","name":"Deep Work","author":"Carl Jung","created_at":0}],"error":null}`, responseBody)
		})
	})
}

func TestBooksHandler_GetBook(t *testing.T) {
	t.Run("invalid url", func(t *testing.T) {
		it(func() {
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
		it(func() {
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
		it(func() {
			invalidUid := "a-valid-uuid"
			url := fmt.Sprintf("/%s/", invalidUid)
			req, err := http.NewRequest(http.MethodGet, url, nil)
			if err != nil {
				t.Fatal(err)
			}
			recorder := httptest.NewRecorder()
			mockDatabase.On("GetBook", invalidUid).Return(nil)

			subject.router.ServeHTTP(recorder, req)

			expectedResponse := fmt.Sprintf(`{"data":null,"error":"%s"}`, contants.ErrItemNotFound)
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
