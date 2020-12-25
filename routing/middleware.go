package routing

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func ContentTypeMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(writer, request)
	})
}

func LoggingMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		log.Printf("request: %#v - %#v", request.RequestURI, mux.Vars(request))
		next.ServeHTTP(writer, request)
	})
}
