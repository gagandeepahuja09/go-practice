package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/stretchr/gomniauth"
	"github.com/stretchr/objx"
)

type authHandler struct {
	next http.Handler
}

func (h *authHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if cookie, err := r.Cookie("auth"); err == http.ErrNoCookie || cookie.Value == "" {
		// redirect to login page in this case
		w.Header().Set("Location", "/login")
		w.WriteHeader(http.StatusTemporaryRedirect)
	} else if err != nil {
		panic(err.Error())
	} else {
		h.next.ServeHTTP(w, r)
	}
}

// loginHandler handles the 3rd party login process.
// format: /auth/{action}/{provider}
// we have preferred a functional approach instead of creating a struct
// because we don't expect it to store any state.
func loginHandler(w http.ResponseWriter, r *http.Request) {
	segs := strings.Split(r.URL.Path, "/")
	action := segs[2]
	provider := segs[3]
	switch action {
	case "login":
		provider, err := gomniauth.Provider(provider)
		if err != nil {
			log.Fatalln("Error when trying to get provider", provider, "-", err)
		}
		loginUrl, err := provider.GetBeginAuthURL(nil, nil)
		if err != nil {
			log.Fatalln("Error when trying to getBeginAuthURL for provider", provider, "-", err)
		}
		w.Header().Set("Location", loginUrl)
		w.WriteHeader(http.StatusTemporaryRedirect)
	case "callback":
		provider, err := gomniauth.Provider(provider)
		if err != nil {
			log.Fatalln("Error when trying to get provider", provider, "-", err)
		}
		creds, err := provider.CompleteAuth(objx.MustFromURLQuery(r.URL.RawQuery))
		if err != nil {
			log.Fatalln("Error when trying to complete auth for provider", provider, "-", err)
		}
		user, err := provider.GetUser(creds)
		if err != nil {
			log.Fatalln("Error when trying to get user from ", provider, "-", err)
		}

		m := md5.New()
		io.WriteString(m, strings.ToLower(user.Email()))
		userId := fmt.Sprintf("%x", m.Sum(nil))

		authCookieValue := objx.New(map[string]interface{}{
			"name":       user.Name(),
			"avatar_url": user.AvatarURL(),
			"email":      user.Email(),
			"userid":     userId,
		}).MustBase64()

		http.SetCookie(w, &http.Cookie{
			Name:  "auth",
			Value: authCookieValue,
			Path:  "/",
		})

		// storing non-signed cookies like this is fine for incidental information such as user name
		// however we should avoid storing any sensitive information using non-signed cookies as it's
		// easy for people to access and change that data.

		w.Header()["Location"] = []string{"/chat"}
		w.WriteHeader(http.StatusTemporaryRedirect)

	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Auth action %s not supported", action)
	}
}

func MustAuth(handler http.Handler) http.Handler {
	return &authHandler{next: handler}
}
