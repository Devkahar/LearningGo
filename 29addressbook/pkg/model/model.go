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
	DOB         string   `db:"dob" json:"dob" validate:"required"`
	Address     *Address `json:"address" gorm:"foreignKey:ID" validate:"required"`
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
	Country       string `json:"country"`
	State         string `json:"state"`
	City          string `json:"city"`
	ZipCode       string `json:"zipCode"`
	StreetAddress string `json:"streetAddress"`
}
type Message struct {
	Str string `json:"message"`
}

type APIResponse struct {
	Status  int         `json:"status"`
	Err     error       `json:"error"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

var ErrorMessages = map[string]string{
	"Address": "address is required",
	"DOB":     "date of birth is Required",
}
