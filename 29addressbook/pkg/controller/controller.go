package controller

import (
	"addressbook/pkg/db"
	"addressbook/pkg/model"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

// Home Page
func ServeHome(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	fmt.Println("Home Page")
	w.Write([]byte("<h1>Address Book</h1>"))
}

// List of contact
func ContactList(w http.ResponseWriter, r *http.Request) (res model.APIResponse) {
	defer r.Body.Close()
	fmt.Println("Get Contact List")
	contacts, err := getContactList()
	fmt.Println("Error -> ", err)
	if err != nil {
		res.Err = err
		res.Status = http.StatusInternalServerError
		return
	}
	res.Data = contacts
	res.Message = "list of contact"
	res.Status = http.StatusOK
	return
}

// add contact to list
func AddContact(w http.ResponseWriter, r *http.Request) (res model.APIResponse) {
	defer func() {
		if res.Err == nil {
			res.Message = "Contact Added Successfully."
			res.Status = http.StatusCreated
		}
		r.Body.Close()
	}()
	fmt.Println("Add Contact")
	var contact model.Contact
	if res.Err = json.NewDecoder(r.Body).Decode(&contact); res.Err != nil {
		return
	}
	// fmt.Println("Address-> city", contact.Address.City)
	validate := validator.New()
	res.Err = validate.Struct(contact)
	if res.Err != nil {
		// this check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		// value most including myself do not usually have code like this.
		if _, ok := res.Err.(*validator.InvalidValidationError); ok {
			fmt.Println(res.Err)
			return
		}
		var fieldErrors []string
		for _, err := range res.Err.(validator.ValidationErrors) {
			// fmt.Println(err.Namespace())
			// fmt.Println(err.Field())
			// fmt.Println(err.StructNamespace())
			// fmt.Println(err.StructField())
			// fmt.Println(err.Tag())
			// fmt.Println(err.ActualTag())
			// fmt.Println(err.Kind())
			// fmt.Println(err.Type())
			// fmt.Println(err.Value())
			// fmt.Println(err.Param())
			// fmt.Println()
			field := err.Field()
			if message, ok := model.ErrorMessages[field]; ok {
				fieldErrors = append(fieldErrors, message)
			} else {
				fieldErrors = append(fieldErrors, err.Error())
			}
			res.Data = fieldErrors
		}
		res.Err = errors.New("invalid fields")
		return
	}
	if res.Err = insertContact(contact); res.Err != nil {
		return
	}
	return
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
		// defer helper.HandelApiError(w, &model.HttpError{Status: 400, Message: err})
		return
	}
	updateContact(id, &contact)
	success := model.Message{Str: "Contact Updated successfully"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(success)
}

func FileUpload(w http.ResponseWriter, r *http.Request) (res model.APIResponse) {
	r.ParseMultipartForm(10 * 1024 * 1024) // 10 MB file allowed
	file, handler, err := r.FormFile("jsonFile")
	if err != nil {
		return model.APIResponse{
			Status: http.StatusBadRequest,
			Err:    err,
		}
	}
	defer file.Close()
	// get file type
	fileType := handler.Header.Get("Content-Type")
	fmt.Println("File type", fileType)
	if fileType != "application/json" {
		res.Status = http.StatusBadRequest
		res.Err = fmt.Errorf("invalid file, file must be json")
		return
	}
	dataBytes, _ := ioutil.ReadAll(file)
	fmt.Println(string(dataBytes))
	res.Status = http.StatusCreated
	res.Message = "File Readed Successfully"
	return
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

	// err := db.DB.Preload(clause.Associations).Find(&contacts).Error
	err := db.DB.Joins("Contact").Joins("Name").Joins("Phone").Joins("Address").Find(&contacts).Error
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
