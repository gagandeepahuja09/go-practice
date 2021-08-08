package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	call("localhost:8080/basicAuth", "POST")
}

func call(url string, method string) error {
	// create client
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	// create request
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return fmt.Errorf("Got Error %s", err)
	}
	// set the basic auth
	req.SetBasicAuth("abc", "122")
	// making request using client.do
	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Got error %s", err)
	}
	defer res.Body.Close()
	return nil
}
