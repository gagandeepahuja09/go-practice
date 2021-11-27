package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"path"
)

func uploadHandler(w http.ResponseWriter, req *http.Request) {
	userId := req.FormValue("userid")

	fmt.Println(userId)

	// multipart.file is also an io.Reader
	file, header, err := req.FormFile("avatarFile")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	fileName := path.Join("avatars", userId+path.Ext(header.Filename))
	err = ioutil.WriteFile(fileName, data, 0777)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(w, "Successful")
}
