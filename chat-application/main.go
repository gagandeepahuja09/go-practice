package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"text/template"

	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/google"
	"github.com/stretchr/objx"
	"go-practice.com/chat-application/trace"
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

	// set the UserData so that it can be used in templates.
	data := map[string]interface{}{
		"Host": r.Host,
	}

	if authCookie, err := r.Cookie("auth"); err == nil {
		data["UserData"] = objx.MustFromBase64(authCookie.Value)
	}
	// pass the details of the request as data
	// this tells the template to render itself using data that can be extract from
	// http.Request which happens to include the host address.
	t.templ.Execute(w, data)
}

func newRoom(avatar Avatar) *room {
	return &room{
		forward: make(chan *message),
		join:    make(chan *client),
		leave:   make(chan *client),
		clients: make(map[*client]bool),
		// by default will be created with a nil tracer
		tracer: trace.Off(),
		avatar: avatar,
	}
}

func setUpGomniauth() {
	// sends state data b/w the client & server along with a signature checksum.
	// checksum ensures that the state values haven't been tampered with while transmiting.
	// the security key is used for creating the hash in a way that it is impossible
	// to recreate the same hash.
	gomniauth.SetSecurityKey("there_goes_name_of_mine_gagandeep_singh_ahuja_some_long_key_here_goes_here")
	gomniauth.WithProviders(
		google.New("321365375874-75ehdnd9f0128st7tdraqnr552uicl47.apps.googleusercontent.com",
			"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",
			"http://localhost:8080/auth/callback/google"),
	)
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:   "auth",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	})

	w.Header().Set("Location", "/chat")
	w.WriteHeader(http.StatusTemporaryRedirect)
}

func main() {
	var addr = flag.String("addr", ":8080", "The address of the application")
	flag.Parse()

	setUpGomniauth()

	// we don't have to create an instance of authAvatar, so no memory was
	// allocated.
	// Imagine the size of the potential savings if our room had 1000s of rooms.
	r := newRoom(UseGravatar)
	r.tracer = trace.New(os.Stdout)

	http.Handle("/chat", MustAuth(&templateHandler{fileName: "chat.html"}))
	http.Handle("/login", &templateHandler{fileName: "login.html"})
	http.Handle("/upload", &templateHandler{fileName: "upload.html"})
	http.Handle("/room", r)
	http.HandleFunc("/auth/", loginHandler)
	http.HandleFunc("/logout", logoutHandler)
	http.HandleFunc("/uploader", uploadHandler)

	// we are running the room in a separate goroutine so that the chatting operation
	// occurs in the background, allowing our main thread to run the web server.
	go r.run()

	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("Listen and serve:", err)
	}
}
