package rest

import (
	"context"
	"crypto/rand"
	"fmt"
	"net/http"
	"sync"
	"time"
)

var session = sync.Map{}

func login(w http.ResponseWriter, r *http.Request) {
	user := r.FormValue("user")
	pass := r.FormValue("pass")

	if user == "" || pass == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if user != "a" || pass != "a" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var token string
	for _, found := session.Load(token); found || token == ""; {
		b := make([]byte, 32)
		rand.Read(b)
		token = fmt.Sprintf("%x", b)
	}
	session.Store(token, time.Now())

	w.Header().Set("x-access-token", token)

	fmt.Println("login", user, pass)
}

func auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.RequestURI)

		var token = r.Header.Get("x-access-token")

		value, found := session.Load(token)
		if !found {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		ctx := context.WithValue(r.Context(), "session", value)
		w.Header().Add("Expires", "Mon, 26 Jul 1997 05:00:00 GMT")
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func authp(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.RequestURI, r.Method, r.URL.Query())
		w.Header().Add("Expires", "Mon, 26 Jul 1997 05:00:00 GMT")
		next.ServeHTTP(w, r)
	})
}
