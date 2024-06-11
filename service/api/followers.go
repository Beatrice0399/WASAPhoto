package api

import (
	"net/http"
	"strconv"

	"github.com/Beatrice0399/WASAPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

// function that allows to follow a user
func (rt *_router) followUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	string_fid := ps.ByName("fid")
	myId, err := rt.get_uid_path(ps)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	requestingUserId := extractBearer(r.Header.Get("Authorization"))

	// users can't follow themselves
	if requestingUserId == ps.ByName("uid") {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if ps.ByName("fid") != requestingUserId {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fid, err := strconv.Atoi(string_fid)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = rt.db.FollowUser(fid, myId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)

}

// Function that allows to unfollow a user
func (rt *_router) unfollowUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	string_fid := ps.ByName("fid")

	myId, err := rt.get_uid_path(ps)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	requestingUserId := extractBearer(r.Header.Get("Authorization"))

	// users can't unfollow themselves
	if requestingUserId == ps.ByName("uid") {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if ps.ByName("fid") != requestingUserId {
		w.WriteHeader(http.StatusBadRequest)

		return
	}
	fid, err := strconv.Atoi(string_fid)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = rt.db.UnfollowUser(fid, myId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
