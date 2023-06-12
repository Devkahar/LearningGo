package main

import (
	"addressbook/pkg/db"
	"addressbook/pkg/router"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var DB *gorm.DB

func main() {
	fmt.Println("Address Book Application")
	db.CreateConnection()
	r := mux.NewRouter()
	router.RegisterRoutes(r)

	http.ListenAndServe(":4000", r)
}
