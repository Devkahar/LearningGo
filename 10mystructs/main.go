package main

import "fmt"

func main() {
	fmt.Println("Struct in golang")
	// No Inheritance in golang; No Super or parent
	dev := User{"Dev", "dev@go.dev", true, 21}
	fmt.Println(dev)
	fmt.Printf("Dev , %+v", dev)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int64
}
