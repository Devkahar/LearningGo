package controller

import (
	"addressbook/pkg/db"
	"addressbook/pkg/model"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// Home Page
func ServeHome(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Home Page")
	w.Write([]byte("<h1>Address Book</h1>"))
}

// List of contact
func ContactList(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get Contact List")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(db.Contacts)

	defer r.Body.Close()
}

// add contact to list
func AddContact(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Add Contact")
	var contact model.Contact
	json.NewDecoder(r.Body).Decode(&contact)
	contact.Id = uuid.NewString()
	db.Contacts = append(db.Contacts, contact)
	success := model.Message{Str: "Contact Added Successfully."}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(success)
	defer r.Body.Close()
}

// add contact to list
func RemoveContact(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Remove Contact")
	params := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	for idx, contact := range db.Contacts {
		if contact.Id == params["id"] {
			db.Contacts = append(db.Contacts[:idx], db.Contacts[idx+1:]...)
			success := model.Message{Str: "Contact Deleted Successfully."}
			json.NewEncoder(w).Encode(success)
			return
		}
	}
	w.WriteHeader(http.StatusBadRequest)
	errorMessage := model.Message{Str: "Contact Not Found."}
	json.NewEncoder(w).Encode(errorMessage)
	defer r.Body.Close()
}

// add contact to list
func ModifyContact(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Modify Contact")
	params := mux.Vars(r)
	var newContact model.Contact
	json.NewDecoder(r.Body).Decode(&newContact)
	w.Header().Set("Content-Type", "application/json")
	for idx, contact := range db.Contacts {
		if contact.Id == params["id"] {
			db.Contacts = append(db.Contacts[:idx], db.Contacts[idx+1:]...)
			db.Contacts = append(db.Contacts, newContact)
			success := model.Message{Str: "Contact Updated Successfully."}
			json.NewEncoder(w).Encode(success)
			return
		}
	}
	w.WriteHeader(http.StatusBadRequest)
	errorMessage := model.Message{Str: "Contact Not Found."}
	json.NewEncoder(w).Encode(errorMessage)
	defer r.Body.Close()
}
