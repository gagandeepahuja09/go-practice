package main

import "net/http"

//

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/v1/teachers", teacherHandler)

	// studentHandler is a type implementing the handler interface
	sHandler := studentHandler{}
	mux.Handle("/v1/student", sHandler)

	s := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	s.ListenAndServe()
}

// req is a struct
// ResponseWriter is an interface
func teacherHandler(res http.ResponseWriter, _ *http.Request) {
	data := []byte("V1 of teacher is called")
	res.WriteHeader(200)
	res.Write(data)
}

type studentHandler struct{}

func (sh studentHandler) ServeHTTP(res http.ResponseWriter, _ *http.Request) {
	data := []byte("V1 of student is called")
	res.WriteHeader(200)
	res.Write(data)
}
