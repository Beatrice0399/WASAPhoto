package api

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Beatrice0399/WASAPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) banUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	string_bid := rt.get_bid(ps)
	myId, err := rt.get_myid_path(ps)
	log.Printf("1")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	requestingUid := extractBearer(r.Header.Get("Authorization"))
	log.Printf("2")
	// users can't follow themselves
	if requestingUid == string_bid {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Printf("3")
	if ps.ByName("uid") != requestingUid {
		w.WriteHeader(http.StatusBadRequest)

		return
	}
	log.Printf("4")
	bid, err := strconv.Atoi(string_bid)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Printf("5")
	err = rt.db.BanUser(myId, bid)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Printf("6")
	w.WriteHeader(http.StatusNoContent)
}

func (rt *_router) unbanUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	log.Printf("0")
	string_bid := rt.get_bid(ps)
	myId, err := rt.get_myid_path(ps)
	log.Printf("1")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	requestingUid := extractBearer(r.Header.Get("Authorization"))
	log.Printf("2")
	// users can't follow themselves
	if requestingUid == string_bid {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Printf("3")
	if ps.ByName("uid") != requestingUid {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Printf("4")
	bid, err := strconv.Atoi(string_bid)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Printf("5")
	err = rt.db.UnbanUser(myId, bid)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Printf("6")
	w.WriteHeader(http.StatusNoContent)
}
