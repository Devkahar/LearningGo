package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Here we will Listern to cancellation Event Emitted by Client.
func main() {
	fmt.Println("Learning About context")
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")
	http.ListenAndServe(":8000", r)
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	select {
	case <-time.After(2 * time.Second):
		w.Write([]byte("Hey Welcome to context page"))
	case <-ctx.Done():
		fmt.Println("User Left Request in between")
	}

}
