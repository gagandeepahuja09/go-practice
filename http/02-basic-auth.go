package main

import (
	"fmt"
	"net/http"
)

var (
	password = "12"
	username = "abc"
)

func main() {
	handler := http.HandlerFunc(handleRequest)
	http.HandleFunc("/basicAuth", handler)
	http.ListenAndServe(":8080", nil)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	u, p, ok := r.BasicAuth()
	fmt.Println(u, p, ok)
	if !ok {
		fmt.Println("Error parsing basic auth")
		w.WriteHeader(401)
		return
	}
	if u != username || p != password {
		fmt.Printf("Username or password provided is not correct: %s and %s", u, p)
		w.WriteHeader(401)
		return
	}
	w.WriteHeader(200)
	return
}
