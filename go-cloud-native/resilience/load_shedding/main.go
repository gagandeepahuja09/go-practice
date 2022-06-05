package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const MaxQueueDepth = 1000
const CurrentDepth = 1001

func loadSheddingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if CurrentDepth > MaxQueueDepth {
			log.Println("load shedding engaged")

			http.Error(w, "load shedding engaged", http.StatusServiceUnavailable)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	r := mux.NewRouter()

	r.Use(loadSheddingMiddleware)

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world"))
	})

	log.Fatal(http.ListenAndServe(":8080", r))
}
