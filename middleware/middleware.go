package middleware

import (
	muxHandlers "github.com/gorilla/handlers"
	"net/http"
	"os"
)

func ContentTypeMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(writer, request)
	})
}

func LoggingMiddleWare(next http.Handler) http.Handler {
	return muxHandlers.LoggingHandler(os.Stdout, next)
}
