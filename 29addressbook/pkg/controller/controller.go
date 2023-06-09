package controller

import (
	"addressbook/pkg/db"
	"addressbook/pkg/model"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

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
	json.NewEncoder(w).Encode(contacts)
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
	deleteContact(id)
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

func insertContact(contact model.Contact) error {
	database := db.CreateConnection()
	contactQuery := "INSERT INTO contact(email,dob) VALUES($1,$2) RETURNING id"
	var id int
	err := database.QueryRow(contactQuery, contact.Email, contact.DOB).Scan(&id)
	if err != nil {
		fmt.Println("Panic in contactQuery")
		return err
	}
	nameQuery := "INSERT INTO name(id,firstname,lastname) VALUES($1,$2,$3)"
	err = database.QueryRow(nameQuery, id, contact.Name.FirstName, contact.Name.LastName).Err()
	if err != nil {
		fmt.Println("Panic in nameQuery")
		return err
	}
	phoneQuery := "INSERT INTO phone(id,ptype,pnumber) VALUES($1,$2,$3)"
	err = database.QueryRow(phoneQuery, id, contact.PhoneNumber.PhoneType, contact.PhoneNumber.Number).Err()
	if err != nil {
		fmt.Println("Panic in phoneQuery")
		return err
	}
	addressQuery := "INSERT INTO address(id,country,state,city,zipcode) VALUES($1,$2,$3,$4,$5)"
	err = database.QueryRow(addressQuery, id, contact.Address.Country, contact.Address.State, contact.Address.City, contact.Address.ZipCode).Err()
	if err != nil {
		fmt.Println("Panic in addressQuery")
		return err
	}
	return nil
}

func getContactList() ([]model.Contact, error) {
	var contacts []model.Contact
	contactQuery := "SELECT contact,address,phone,name from contact,address,phone,name where contact.id=address.id AND contact.id=phone.id AND contact.id=name.id;"
	database := db.CreateConnection()
	rows, err := database.Query(contactQuery)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		var contact model.Contact
		var name model.Name
		var address model.Address
		var phone model.Phone
		err = rows.Scan(&contact, &address, &phone, &name)
		if err != nil {
			fmt.Println("Error while fetching contactList", err)
		}
		contact.Name = &name
		contact.Address = &address
		contact.PhoneNumber = &phone
		contacts = append(contacts, contact)
	}
	fmt.Println(contacts)
	return contacts, nil
}

// update contact
func updateContact(id int, contact *model.Contact) error {
	database := db.CreateConnection()
	sqlQuery := "UPDATE contact SET email=$2, dob=$3 WHERE id=$1"
	_, err := database.Exec(sqlQuery, contact.Id, contact.Email, contact.DOB)
	if err != nil {
		log.Fatalf("Fail to update in DB %v", err)
	}
	sqlQuery = "UPDATE name SET firstname=$2, lastname=$3 WHERE id=$1"
	_, err = database.Exec(sqlQuery, contact.Id, contact.Name.FirstName, contact.Name.LastName)
	if err != nil {
		log.Fatalf("Fail to update in DB %v", err)
	}
	sqlQuery = "UPDATE address SET country=$2, state=$3, city=$4, zipcode=$5 WHERE id=$1"
	_, err = database.Exec(sqlQuery, contact.Id, contact.Address.Country, contact.Address.State, contact.Address.City, contact.Address.ZipCode)
	if err != nil {
		log.Fatalf("Fail to update in DB %v", err)
	}
	sqlQuery = "UPDATE phone SET ptype=$2, pnumber=$3 WHERE id=$1;"
	res, err := database.Exec(sqlQuery, contact.Id, contact.PhoneNumber.PhoneType, contact.PhoneNumber.Number)
	if err != nil {
		log.Fatalf("Fail to update in DB %v", err)
	}
	rowsAffectes, err := res.RowsAffected()
	if err != nil {
		return err
	}
	fmt.Println(rowsAffectes)
	return nil
}

// Delete Contact
func deleteContact(id int) error {
	sqlQuery := "DELETE FROM contact WHERE id=$1"
	database := db.CreateConnection()
	err := database.QueryRow(sqlQuery, id)
	if err != nil {
		fmt.Println(err)
	}
	defer database.Close()
	return nil
}
