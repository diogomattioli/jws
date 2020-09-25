package main

import (
	"context"
	"crypto/rand"
	"fmt"
	"jws/system"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "It's working!")
}

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

func main() {
	system.InitDB()

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", home)
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("."))))
	router.HandleFunc("/login", login).Methods("POST")

	sub := router.PathPrefix("/auth").Subrouter()
	sub.Use(authp)
	sub.HandleFunc("/{type}/", system.View).Methods("GET")
	sub.HandleFunc("/{type}/", system.Create).Methods("POST")
	sub.HandleFunc("/{type}/{id}", system.Retrieve).Methods("GET")
	sub.HandleFunc("/{type}/", system.Update).Methods("PATCH")
	sub.HandleFunc("/{type}/{id}", system.Delete).Methods("DELETE")
	sub.HandleFunc("/{type}/{id}/lock", system.Lock).Methods("GET")
	sub.HandleFunc("/{type}/{id}/unlock", system.Unlock).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
