package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getMyStream(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	/*myId, err := rt.getMyId(r)
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
	stream, err := rt.db.GetMyStream(myid)
	if err != nil {
		rt.responsError(http.StatusBadRequest, err.Error(), w)
		return
	}
	rt.responseJson(stream, w)
}
