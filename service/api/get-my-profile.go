package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getMyProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	/*
		myId, err := rt.getMyId(r)
		if err != nil {
			rt.responsError(http.StatusBadRequest, err.Error(), w)
			return
		}
	*/
	myid, err := rt.get_myid_path(ps)
	if err != nil {
		rt.responsError(http.StatusBadRequest, err.Error(), w)
		return
	}
	profile, err := rt.db.GetMyProfile(myid)
	if err != nil {
		rt.responsError(http.StatusBadRequest, err.Error(), w)
		return
	}
	rt.responseJson(profile, w)

}
