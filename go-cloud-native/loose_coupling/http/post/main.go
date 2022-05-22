package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

const json = `{ "name": "Gagan", "age": 23 }`

func main() {
	in := strings.NewReader(json)

	resp, err := http.Post("http://example.com/upload", "text/json", in)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	message, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(message))
}
