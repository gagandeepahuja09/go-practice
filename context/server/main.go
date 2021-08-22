package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// we can use the req.Context() and when we cancel the request from
// the client, it also gets cancelled from the server using the channel
// ctx.Done. As the channel will be closed when we cancel the request.

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Printf("handler started")
	defer log.Printf("handler ended")

	select {
	case <-time.After(10 * time.Second):
		fmt.Fprintln(w, "hello")
	case <-ctx.Done():
		err := ctx.Err()
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8888", nil))
}
