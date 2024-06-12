package api

import (
	"net/http"
	"strconv"

	"github.com/Beatrice0399/WASAPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

// function that allows to ban an user
func (rt *_router) banUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	string_bid := ps.ByName("bid")
	uid, err := rt.get_uid_path(ps)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	requestingUid := extractBearer(r.Header.Get("Authorization"))
	// users can't ban themselves
	if requestingUid == ps.ByName("uid") {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if string_bid != requestingUid {
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	bid, err := strconv.Atoi(string_bid)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = rt.db.BanUser(bid, uid)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// function that allows to unban an user+++
func (rt *_router) unbanUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	string_bid := ps.ByName("bid")
	uid, err := rt.get_uid_path(ps)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	requestingUid := extractBearer(r.Header.Get("Authorization"))

	// users can't unban themselves
	if requestingUid == ps.ByName("uid") {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if string_bid != requestingUid {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	bid, err := strconv.Atoi(string_bid)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = rt.db.UnbanUser(bid, uid)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
