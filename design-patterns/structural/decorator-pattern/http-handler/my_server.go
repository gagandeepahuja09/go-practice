package main

import (
	"fmt"
	"net/http"
)

type MyServer struct {
}

func (m *MyServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello Decorator!")
}
