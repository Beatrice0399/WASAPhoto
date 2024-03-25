package api

import (
	"encoding/json"
	"net/http"
)

func (rt _router) responseJson(body interface{}, w http.ResponseWriter) error {
	w.Header().Set("Content-type", "application/json")
	bytes, err := json.MarshalIndent(body, "", "\t")
	if err != nil {
		return err
	}
	w.Write(bytes)
	return nil
}

func (rt _router) responsError(code int, body interface{}, w http.ResponseWriter) error {
	w.WriteHeader(code)
	w.Header().Set("Content-type", "application/json")
	bytes, err := json.MarshalIndent(body, "", "\t")
	if err != nil {
		return err
	}
	w.Write(bytes)
	return nil
}
