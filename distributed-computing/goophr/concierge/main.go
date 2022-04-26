package main

import "net/http"

// import "github.com/stretchr/gomniauth/common"

func main() {
	common.Log("Adding API handlers")
	http.HandleFunc("/api/feeder", api.FeedHandler)

	api.StartFeederSystem()
}
