package main

import (
	"addressbook/pkg/router"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Address Book Application")
	r := mux.NewRouter()
	router.RegisterRoutes(r)
	http.ListenAndServe(":4000", r)
}
