package router

import (
	"addressbook/pkg/controller"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {

	r.HandleFunc("/", controller.ServeHome).Methods("GET")
	r.HandleFunc("/contacts", controller.ContactList).Methods("GET")
	r.Handle("/contact", handlerFunc(controller.AddContact)).Methods("POST")
	r.HandleFunc("/contact/{id}", controller.RemoveContact).Methods("DELETE")
	r.HandleFunc("/contact/{id}", controller.ModifyContact).Methods("PUT")
}

type handlerFunc func(w http.ResponseWriter, r *http.Request) error

func (f handlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := f(w, r); err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
	}
}
