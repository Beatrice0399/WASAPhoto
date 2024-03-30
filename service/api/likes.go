package api

import (
	"net/http"
	"strconv"

	"github.com/Beatrice0399/WASAPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) likePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	token := extractBearer(r.Header.Get("Authorization"))
	if isNotLogged(token) {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	myID, err := strconv.Atoi(token)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	uid, err := rt.get_uid_path(ps)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if rt.db.IsBanned(myID, uid) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	lid, err := rt.get_like_id(ps)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if myID != lid {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	phid, err := rt.getPhid(ps)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = rt.db.LikePhoto(phid, myID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (rt *_router) unlikePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	token := extractBearer(r.Header.Get("Authorization"))
	if isNotLogged(token) {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	myID, err := strconv.Atoi(token)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	uid, err := rt.get_uid_path(ps)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if rt.db.IsBanned(myID, uid) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	like_id, err := rt.get_like_id(ps)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if myID != like_id {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	phid, err := rt.getPhid(ps)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = rt.db.UnlikePhoto(phid, myID, like_id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
