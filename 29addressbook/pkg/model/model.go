package model

type Contact struct {
	Id          string   `json:"id"`
	Name        *Name    `json:"name"`
	Email       string   `json:"email"`
	PhoneNumber string   `json:"phonenumber"`
	DOB         *Date    `json:"dob"`
	Address     *Address `json:"address"`
	Img         *Img     `json:"image"`
}
type Name struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
type Date struct {
	Day   int `json:"day"`
	Month int `json:"month"`
	Year  int `json:"year"`
}
type Address struct {
	Country       string `json:"country"`
	State         string `json:"firstName"`
	City          string `json:"city"`
	ZipCode       string `json:"zipCode"`
	StreetAddress string `json:"streetAddress"`
}
type Img struct {
	Url string `json:"url"`
}

type Message struct {
	Str string `json:"message"`
}
