package router

import (
	"addressbook/pkg/controller"
	"addressbook/pkg/model"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/", controller.ServeHome).Methods("GET")
	r.Handle("/contacts", handlerFunc(controller.ContactList)).Methods("GET")
	r.Handle("/contact", handlerFunc(controller.AddContact)).Methods("POST")
	r.HandleFunc("/contact/{id}", controller.RemoveContact).Methods("DELETE")
	r.HandleFunc("/contact/{id}", controller.ModifyContact).Methods("PUT")
	r.Handle("/upload", handlerFunc(controller.FileUpload)).Methods("POST")
}

type handlerFunc func(w http.ResponseWriter, r *http.Request) model.APIResponse

func (f handlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("Recover from panic err: %v", err)
			writeJSON(w, http.StatusInternalServerError, model.APIResponse{
				Err:    fmt.Errorf("%v", err),
				Status: http.StatusInternalServerError,
			})
		}
	}()
	res := f(w, r)

	if res.Err != nil {
		if res.Status == 0 {
			res.Status = http.StatusBadRequest
		}
		// Here some error cases can be written to send different error messages.
		// http.Error(w, res.Err.Error(), res.Status)
		// return
	}
	fmt.Println("Above to call success write json")
	if err := writeJSON(w, res.Status, res); err != nil {
		fmt.Println("Complete")
		writeJSON(w, http.StatusInternalServerError, model.APIResponse{
			Err:    errors.New("internal server error"),
			Status: http.StatusInternalServerError,
		})
	}
}

func writeJSON(w http.ResponseWriter, status int, res model.APIResponse) error {
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(res)
}
