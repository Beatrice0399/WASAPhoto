package api

import (
	"net/http"
	"strconv"

	"github.com/Beatrice0399/WASAPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) followUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	string_fid := rt.get_fid(ps)
	myId, err := rt.get_myid_path(ps)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	requestingUserId := extractBearer(r.Header.Get("Authorization"))

	// users can't follow themselves
	if requestingUserId == string_fid {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if ps.ByName("uid") != requestingUserId {
		w.WriteHeader(http.StatusBadRequest)

		return
	}
	fid, err := strconv.Atoi(string_fid)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = rt.db.FollowUser(myId, fid)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)

}

func (rt *_router) unfollowUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	string_fid := rt.get_fid(ps)

	myId, err := rt.get_myid_path(ps)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	requestingUserId := extractBearer(r.Header.Get("Authorization"))

	// users can't follow themselves
	if requestingUserId == string_fid {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if ps.ByName("uid") != requestingUserId {
		w.WriteHeader(http.StatusBadRequest)

		return
	}
	fid, err := strconv.Atoi(string_fid)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = rt.db.UnfollowUser(myId, fid)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
