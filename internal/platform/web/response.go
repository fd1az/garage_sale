package web

import (
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

func Respond(w http.ResponseWriter, val interface{}, statusCode int) error {
	data, err := json.Marshal(val)
	if err != nil {
		return errors.Wrapf(err, "marshaling value to json")
	}
	w.Header().Set("content-type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)

	if _, err := w.Write(data); err != nil {
		return errors.Wrapf(err, "wrting to client")
	}
	return nil
}

//RespondError knows how to handle error going to the client
func RespondError(w http.ResponseWriter, err error) error {

	if webErr, ok := err.(*Error); ok {
		resp := ErrorResponse{
			Error: webErr.Err.Error(),
		}
		return Respond(w, resp, webErr.Status)
	}
	resp := ErrorResponse{
		Error: http.StatusText(http.StatusInternalServerError),
	}
	return Respond(w, resp, http.StatusInternalServerError)
}
