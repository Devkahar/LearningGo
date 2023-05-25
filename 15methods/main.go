package main

import "fmt"

func main() {
	fmt.Println("Learning Methods in Go")
	// no inheritance in golang; No super or parent
	dev := User{"Dev", "dev@gmail.com", true, 20}
	fmt.Printf("Dev Details %+v", dev)
	dev.GetStatus()
	dev.NewEmail()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() bool {
	fmt.Println("User status ", u.Status)
	return u.Status
}

func (u User) NewEmail() {
	u.Email = "dev@go.dev"
	fmt.Println("User Email Address ", u.Email)
}
