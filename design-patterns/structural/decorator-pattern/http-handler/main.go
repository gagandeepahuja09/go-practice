package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/", &MyServer{})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
