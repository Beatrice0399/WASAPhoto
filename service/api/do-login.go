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
		rt.responsError(http.StatusBadRequest, "Invalid username login", w)
		return
	}
	id, err := rt.db.DoLogin(username)
	if err != nil {
		rt.responsError(http.StatusBadRequest, err.Error(), w)
		return
	}
	rt.responseJson(id, w)
}
