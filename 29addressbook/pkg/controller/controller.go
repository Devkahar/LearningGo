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
	defer r.Body.Close()
	fmt.Println("Home Page")
	w.Write([]byte("<h1>Address Book</h1>"))
}

// List of contact
func ContactList(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	fmt.Println("Get Contact List")
	w.Header().Set("Content-Type", "application/json")
	contacts, err := getContactList()
	if err != nil {
		fmt.Println(err)
	}
	json.NewEncoder(w).Encode(&contacts)
}

// func checkError(w http.ResponseWriter) {
// 	if err := recover(); err != nil {
// 		helper.HandelApiError(w, &model.HttpError{Status: 400, Message: fmt.Sprintf("%v", err)})
// 	}
// }

// add contact to list
func AddContact(w http.ResponseWriter, r *http.Request) error {
	defer r.Body.Close()
	// defer checkError(w)
	fmt.Println("Add Contact")
	var contact model.Contact
	err := json.NewDecoder(r.Body).Decode(&contact)
	if err != nil {
		return err
	}
	contact.Address = nil
	err = insertContact(contact)
	if err != nil {
		return err
	}
	success := model.Message{Str: "Contact Added Successfully."}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(success)
	return nil
}

// remove contact form list
func RemoveContact(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	fmt.Println("Remove Contact")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	err := deleteContact(id)
	if err != nil {
		panic(err)
	}
	success := model.Message{Str: "Contact Deleted successfully"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(success)
}

// modify contact
func ModifyContact(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	fmt.Println("Modify Contact")
	params := mux.Vars(r)
	var contact model.Contact
	json.NewDecoder(r.Body).Decode(&contact)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		// defer helper.HandelApiError(w, &model.HttpError{Status: 400, Message: err.Error()})
		return
	}
	updateContact(id, &contact)
	success := model.Message{Str: "Contact Updated successfully"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(success)
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
