package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	lw := logWriter{}
	io.Copy(lw, resp.Body)
}

func (lw logWriter) Write(b []byte) (int, error) {
	fmt.Println(string(b))
	fmt.Println("No. of bytes scanned = ", len(b))
	return len(b), nil
}
