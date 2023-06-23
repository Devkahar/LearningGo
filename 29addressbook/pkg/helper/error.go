package helper

import (
	"addressbook/pkg/model"
	"encoding/json"
	"net/http"
)

func HandelApiError(w http.ResponseWriter, httpErr *model.HttpError) {
	w.WriteHeader(httpErr.Status)
	data, err := json.Marshal(httpErr)
	if err != nil {
		httpErr.Message = "Someting went wrong"
	}
	w.Write([]byte(data))
}
