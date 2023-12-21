package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) setMyUsername(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var username string
	err := json.NewDecoder(r.Body).Decode(&username)
	if err != nil {
		rt.responsError(http.StatusBadRequest, err.Error(), w)
		return
	}
	myId, err := rt.get_myid_path(ps)
	if err != nil {
		rt.responsError(http.StatusBadRequest, err.Error(), w)
		return
	}
	name, err := rt.db.SetMyUsername(myId, username)

	if err != nil {
		rt.responsError(http.StatusBadRequest, err.Error(), w)
		return
	}
	rt.responseJson(name, w)
}
