package api

import (
	"net/http"
	"strconv"

	"github.com/Beatrice0399/WASAPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) banUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	string_bid := rt.get_bid(ps)
	myId, err := rt.get_uid_path(ps)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	requestingUid := extractBearer(r.Header.Get("Authorization"))
	// users can't ban themselves
	if requestingUid == string_bid {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if ps.ByName("uid") != requestingUid {
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	bid, err := strconv.Atoi(string_bid)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = rt.db.BanUser(myId, bid)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (rt *_router) unbanUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	string_bid := rt.get_bid(ps)
	myId, err := rt.get_uid_path(ps)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	requestingUid := extractBearer(r.Header.Get("Authorization"))

	// users can't unban themselves
	if requestingUid == string_bid {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if ps.ByName("uid") != requestingUid {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	bid, err := strconv.Atoi(string_bid)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = rt.db.UnbanUser(myId, bid)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
