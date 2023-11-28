package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var username string
	err := json.NewDecoder(r.Body).Decode(&username)
	if err != nil {
		w.Header().Set("Content-type", "application/json")
		w.Write([]byte("Invalid username login"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id, err := rt.db.DoLogin(username)
	w.Header().Set("Content-type", "application/json")
	id_byte, err := json.Marshal(id)
	w.Write(id_byte)
}
