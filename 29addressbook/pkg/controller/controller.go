package controller

import (
	"addressbook/pkg/db"
	"addressbook/pkg/model"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"

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
	// Insert 1000 connection
	// for i := 0; i < 1000; i++ {
	// 	err = insertContact(contact)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }
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
func Delay() {
	time.Sleep(100 * time.Millisecond)
}
func executeAddressQuery(contact *model.Contact, mut *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	mut.Lock()
	db.DB.First(&contact.Address, contact.ID)
	mut.Unlock()
	Delay()
}
func executePhoneQuery(contact *model.Contact, mut *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	mut.Lock()
	db.DB.First(&contact.PhoneNumber, contact.ID)
	mut.Unlock()
	Delay()
}
func executeNameQuery(contact *model.Contact, mut *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	mut.Lock()
	db.DB.First(&contact.Name, contact.ID)
	mut.Unlock()
	Delay()
}
func getContactList() ([]model.Contact, error) {
	var contacts []model.Contact
	start := time.Now()

	var mut = &sync.Mutex{}
	var wg = &sync.WaitGroup{}
	err := db.DB.Find(&contacts).Error

	// dbChan := make(chan bool)
	if err != nil {
		return contacts, err
	}
	// dbChan <- true

	// Without coroutines
	for idx := range contacts {
		db.DB.First(&contacts[idx].Address, contacts[idx].ID)
		Delay()
		db.DB.First(&contacts[idx].Name, contacts[idx].ID)
		Delay()
		db.DB.First(&contacts[idx].PhoneNumber, contacts[idx].ID)
		Delay()
	}
	// Querying DB with coroutines

	for idx := range contacts {
		wg.Add(1)
		go executeAddressQuery(&contacts[idx], mut, wg)
		wg.Add(1)
		go executeNameQuery(&contacts[idx], mut, wg)
		wg.Add(1)
		go executePhoneQuery(&contacts[idx], mut, wg)
	}
	wg.Wait()
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
