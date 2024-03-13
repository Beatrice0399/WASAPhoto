package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/Beatrice0399/WASAPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getMyProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
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

func (rt *_router) setMyUsername(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var username string
	err := json.NewDecoder(r.Body).Decode(&username)
	if err != nil {
		ctx.Logger.WithError(err).Error("error decoding json")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	myId, err := rt.get_myid_path(ps)
	if err != nil {
		rt.responsError(http.StatusBadRequest, err.Error(), w)
		return
	}

	valid := validateRequestingUser(strconv.Itoa(myId), extractBearer(r.Header.Get("Authorization")))
	if valid != 0 {
		w.WriteHeader(valid)
		return
	}

	_, err = rt.db.SetMyUsername(myId, username)

	if err != nil {
		ctx.Logger.WithError(err).Error("error executing query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (rt *_router) getUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	username := rt.getUsername(ps)

	uid, err := rt.db.GetId(username)
	if err != nil {
		log.Println("getid ", err)
		rt.responsError(http.StatusBadRequest, err.Error(), w)
		return
	}

	myId, err := rt.get_myid_path(ps)
	if err != nil {
		log.Println("getmyid ", err)
		rt.responsError(http.StatusBadRequest, err.Error(), w)
		return
	}
	profile, err := rt.db.GetUserProfile(uid, myId)
	if err != nil {
		log.Println("getprofile ", err)
		rt.responsError(http.StatusBadRequest, err.Error(), w)
		return
	}
	rt.responseJson(profile, w)

}
