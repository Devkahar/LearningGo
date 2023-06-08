package router

import (
	"addressbook/pkg/controller"

	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/", controller.ServeHome).Methods("GET")
	r.HandleFunc("/contacts", controller.ContactList).Methods("GET")
	r.HandleFunc("/contact", controller.AddContact).Methods("POST")
	r.HandleFunc("/contact/{id}", controller.RemoveContact).Methods("DELETE")
	r.HandleFunc("/contact/{id}", controller.ModifyContact).Methods("PUT")
}
