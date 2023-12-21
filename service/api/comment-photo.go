package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) commentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	phid, err := rt.getPhid(ps)
	if err != nil {
		rt.responsError(http.StatusBadRequest, err.Error(), w)
		return
	}

	myId, err := rt.get_myid_path(ps)
	if err != nil {
		rt.responsError(http.StatusBadRequest, err.Error(), w)
		return
	}

	var msg string
	err = json.NewDecoder(r.Body).Decode(&msg)
	if err != nil {
		rt.responsError(http.StatusBadRequest, err.Error(), w)
		return
	}

	cid, err := rt.db.CommentPhoto(myId, phid, msg)
	if err != nil {
		rt.responsError(http.StatusBadRequest, err.Error(), w)
		return
	}
	rt.responseJson(cid, w)
}
