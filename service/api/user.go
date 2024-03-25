package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/Beatrice0399/WASAPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) setMyUsername(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var new_username string
	err := json.NewDecoder(r.Body).Decode(&new_username)
	if err != nil {
		ctx.Logger.WithError(err).Error("error decoding json")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	uid := ps.ByName("uid")

	valid := validateRequestingUser(uid, extractBearer(r.Header.Get("Authorization")))
	if valid != 0 {
		w.WriteHeader(valid)
		return
	}
	myId, err := strconv.Atoi(uid)
	_, err = rt.db.SetMyUsername(myId, new_username)
	if err != nil {
		ctx.Logger.WithError(err).Error("error executing query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (rt *_router) getUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("Content-Type", "application/json")
	log.Println("prima identifier")
	identifier := extractBearer(r.Header.Get("Authorization"))
	if identifier == "" {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	log.Printf("prima get query")
	username := r.URL.Query().Get("username")
	myId, err := strconv.Atoi(identifier)
	if err != nil {
		ctx.Logger.WithError(err).Error("error get my id")
		w.WriteHeader(http.StatusInternalServerError)
	}
	log.Printf("prima search users")
	users, err := rt.db.SearchUser(myId, username)

	if err != nil {
		ctx.Logger.WithError(err).Error("error database")
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(users)

}

func (rt *_router) getProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("Content-Type", "application/json")
	string_myId := extractBearer(r.Header.Get("Authorization"))
	string_uid := ps.ByName("uid")

	myId, err := strconv.Atoi(string_myId)
	if err != nil {
		ctx.Logger.WithError(err).Error("error get my id")
		w.WriteHeader(http.StatusInternalServerError)
	}

	uid, err := strconv.Atoi(string_uid)
	if err != nil {
		ctx.Logger.WithError(err).Error("error get my id")
		w.WriteHeader(http.StatusInternalServerError)
	}
	profile, err := rt.db.GetUserProfile(uid, myId)
	if err != nil {
		ctx.Logger.WithError(err).Error("error get user profile")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(profile)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("can't create response json")
	}
}
