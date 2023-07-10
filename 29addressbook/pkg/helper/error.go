package helper

import "addressbook/pkg/model"

func ErrorResponse(r *model.APIResponse, statusCode int, err error) {
	r.Status = statusCode
	r.Err = err
}
