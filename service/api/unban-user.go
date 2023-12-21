package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) unbanUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	username := rt.getUsername(ps)

	myId, err := rt.get_myid_path(ps)
	if err != nil {
		rt.responsError(http.StatusBadRequest, err.Error(), w)
		return
	}
	err = rt.db.UnbanUser(myId, username)
	if err != nil {
		rt.responsError(http.StatusBadRequest, err.Error(), w)
		return
	}
	str := "User unbanned"
	rt.responseJson(str, w)
}
