package model

import (
	"fmt"
	"strconv"
	"strings"
)

type Contact struct {
	Id          int64    `db:"id" json:"id"`
	Name        *Name    `json:"name"`
	Email       string   `db:"email" json:"email"`
	PhoneNumber *Phone   `json:"phonenumber,"`
	DOB         string   `db:"dob" json:"dob"`
	Address     *Address `json:"address"`
}
type Name struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
type Phone struct {
	PhoneType string `json:"type"`
	Number    string `json:"number"`
}
type Address struct {
	Country       string `json:"country"`
	State         string `json:"state"`
	City          string `json:"city"`
	ZipCode       string `json:"zipCode"`
	StreetAddress string `json:"streetAddress"`
}
type Message struct {
	Str string `json:"message"`
}

func (contact *Contact) Scan(src interface{}) error {
	// fmt.Println(src)
	// json.Unmarshal(src.([]byte), &c)
	rowStr := string(src.([]byte))
	rowStr = rowStr[1 : len(rowStr)-1]
	data := strings.Split(rowStr, ",")
	contact.Id, _ = strconv.ParseInt(data[0], 0, 64)
	contact.Email = data[1]
	contact.DOB = data[2]
	return nil
}

func (address *Address) Scan(src interface{}) error {
	// fmt.Println(src)
	// json.Unmarshal(src.([]byte), &c)
	rowStr := string(src.([]byte))
	fmt.Println(rowStr)
	rowStr = rowStr[1 : len(rowStr)-1]
	data := strings.Split(rowStr, ",")
	for idx := range data {
		if data[idx] == "\"\"" {
			data[idx] = ""
		}
	}
	address.Country = data[1]
	address.State = data[2]
	address.City = data[3]
	address.ZipCode = data[4]
	return nil
}

func (name *Name) Scan(src interface{}) error {
	// fmt.Println(src)
	// json.Unmarshal(src.([]byte), &c)
	rowStr := string(src.([]byte))
	fmt.Println(rowStr)
	rowStr = rowStr[1 : len(rowStr)-1]
	data := strings.Split(rowStr, ",")
	for idx := range data {
		if data[idx] == "\"\"" {
			data[idx] = ""
		}
	}
	name.FirstName = data[1]
	name.LastName = data[2]
	return nil
}

func (phone *Phone) Scan(src interface{}) error {
	// fmt.Println(src)
	// json.Unmarshal(src.([]byte), &c)
	rowStr := string(src.([]byte))
	fmt.Println(rowStr)
	rowStr = rowStr[1 : len(rowStr)-1]
	data := strings.Split(rowStr, ",")
	for idx := range data {
		if data[idx] == "\"\"" {
			data[idx] = ""
		}
	}
	phone.PhoneType = data[1]
	phone.Number = data[2]
	return nil
}
