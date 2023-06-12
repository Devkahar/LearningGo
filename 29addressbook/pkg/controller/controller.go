package controller

import (
	"addressbook/pkg/db"
	"addressbook/pkg/model"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"gorm.io/gorm/clause"
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
	contacts, err := getContactList()
	if err != nil {
		fmt.Println(err)
	}
	json.NewEncoder(w).Encode(&contacts)
	defer r.Body.Close()
}

// add contact to list
func AddContact(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Add Contact")
	var contact model.Contact
	err := json.NewDecoder(r.Body).Decode(&contact)

	if err != nil {
		panic(err)
	}
	err = insertContact(contact)
	if err != nil {
		panic(err)
	}
	success := model.Message{Str: "Contact Added Successfully."}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(success)
	defer r.Body.Close()
}

// remove contact form list
func RemoveContact(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Remove Contact")
	params := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	id, _ := strconv.Atoi(params["id"])
	err := deleteContact(id)
	if err != nil {
		panic(err)
	}
	success := model.Message{Str: "Contact Deleted successfully"}
	json.NewEncoder(w).Encode(success)
	defer r.Body.Close()
}

// modify contact
func ModifyContact(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Modify Contact")
	params := mux.Vars(r)
	var contact model.Contact
	json.NewDecoder(r.Body).Decode(&contact)
	w.Header().Set("Content-Type", "application/json")
	id, _ := strconv.Atoi(params["id"])
	updateContact(id, &contact)
	success := model.Message{Str: "Contact Updated successfully"}
	json.NewEncoder(w).Encode(success)
	defer r.Body.Close()
}

// DB Operations.
func insertContact(contact model.Contact) error {
	result := db.DB.Table("contacts").Create(&contact)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func getContactList() ([]model.Contact, error) {
	var contacts []model.Contact
	start := time.Now()
	err := db.DB.Preload(clause.Associations).Find(&contacts).Error

	if err != nil {
		return contacts, err
	}
	fmt.Println(time.Since(start))
	return contacts, nil
}

// update contact
func updateContact(id int, contact *model.Contact) error {
	err := db.DB.Save(&contact).Error

	return err
}

// Delete Contact
func deleteContact(id int) error {
	err := db.DB.Delete(&model.Contact{}, id).Error
	return err
}
