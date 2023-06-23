package model

import (
	"gorm.io/gorm"
)

type Contact struct {
	gorm.Model
	ID          uint     `gorm:"primaryKey"`
	Name        *Name    `json:"name" gorm:"foreignKey:ID"`
	Email       string   `db:"email" json:"email"`
	PhoneNumber *Phone   `json:"phonenumber," gorm:"foreignKey:ID"`
	DOB         string   `db:"dob" json:"dob"`
	Address     *Address `json:"address" gorm:"foreignKey:ID"`
}

type Name struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
type Phone struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey"`
	PhoneType string `json:"type"`
	Number    string `json:"number"`
}
type Address struct {
	gorm.Model
	ID            uint   `gorm:"primaryKey"`
	Country       string `json:"country"`
	State         string `json:"state"`
	City          string `json:"city"`
	ZipCode       string `json:"zipCode"`
	StreetAddress string `json:"streetAddress"`
}
type Message struct {
	Str string `json:"message"`
}
type HttpError struct {
	Status  int
	Message string
}

func (Contact) TableName() string {
	return "contacts"
}

func (Name) TableName() string {
	return "names"
}

func (Phone) TableName() string {
	return "phones"
}

func (Address) TableName() string {
	return "addresses"
}
