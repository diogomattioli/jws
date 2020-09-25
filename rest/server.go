package rest

import (
	"fmt"
	"jws/model"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func params(r *http.Request) (interface{}, interface{}, string, int) {
	obj, slice := model.ObjModel(mux.Vars(r)["type"])
	ty := mux.Vars(r)["type"]
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	return obj, slice, ty, id
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "It's working!")
}

func Init() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", home)
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("."))))
	router.HandleFunc("/login", login).Methods("POST")

	sub := router.PathPrefix("/auth").Subrouter()
	sub.Use(authp)
	sub.HandleFunc("/{type}/", list).Methods("GET")
	sub.HandleFunc("/{type}/", create).Methods("POST")
	sub.HandleFunc("/{type}/{id}", retrieve).Methods("GET")
	sub.HandleFunc("/{type}/{id}", update).Methods("PATCH")
	sub.HandleFunc("/{type}/{id}", delete).Methods("DELETE")
	sub.HandleFunc("/{type}/{id}/lock", lock).Methods("GET")
	sub.HandleFunc("/{type}/{id}/unlock", unlock).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
