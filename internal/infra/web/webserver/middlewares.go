package webserver

import (
	"log"
	"net/http"
)

type Middleware struct{}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.Path)

		next.ServeHTTP(w, r)
	})
}
