package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) deletePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	phid, err := rt.getPhid(ps)
	if err != nil {
		rt.responsError(http.StatusBadRequest, err.Error(), w)
		return
	}

	myid, err := rt.get_myid_path(ps)
	if err != nil {
		rt.responsError(http.StatusBadRequest, err.Error(), w)
		return
	}

	err = rt.db.DeletePhoto(phid, myid)
	if err != nil {
		rt.responsError(http.StatusBadRequest, err.Error(), w)
		return
	}
	str := "Photo deleted"
	rt.responseJson(str, w)
}
