package main

import (
	"flag"
	"log"
	"net/http"
	"path/filepath"
	"sync"
	"text/template"
)

type templateHandler struct {
	fileName string
	once     sync.Once
	templ    *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(
			template.ParseFiles(filepath.Join("templates", t.fileName)))
	})
	// pass the details of the request as data
	// this tells the template to render itself using data that can be extract from
	// http.Request which happens to include the host address.
	t.templ.Execute(w, r)
}

func newRoom() *room {
	return &room{
		forward: make(chan []byte),
		join:    make(chan *client),
		leave:   make(chan *client),
		clients: make(map[*client]bool),
	}
}

func main() {
	var addr = flag.String("addr", ":8080", "The address of the application")
	flag.Parse()
	r := newRoom()
	http.Handle("/", &templateHandler{fileName: "chat.html"})
	http.Handle("/room", r)

	// we are running the room in a separate goroutine so that the chatting operation
	// occurs in the background, allowing our main thread to run the web server.
	go r.run()

	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("Listen and serve:", err)
	}
}