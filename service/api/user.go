package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Beatrice0399/WASAPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

// Function that updates an user's username
func (rt *_router) setMyUsername(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	uid := ps.ByName("uid")
	// check the user's identity for the operation
	valid := validateRequestingUser(uid, extractBearer(r.Header.Get("Authorization")))
	if valid != 0 {
		w.WriteHeader(valid)
		return
	}
	myId, err := strconv.Atoi(uid)
	if err != nil {
		ctx.Logger.WithError(err).Error("error converting uid to int")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var new_username Username
	err = json.NewDecoder(r.Body).Decode(&new_username)

	if !validStringUsername(new_username.Username) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if rt.db.UsernameExist(new_username.Username) {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	if err != nil {
		ctx.Logger.WithError(err).Error("error decoding json")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = rt.db.SetMyUsername(myId, new_username.Username)
	if err != nil {
		ctx.Logger.WithError(err).Error("error executing query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

}

// function that retrives all the users whose username contains the given one
func (rt *_router) searchProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("Content-Type", "application/json")
	identifier := extractBearer(r.Header.Get("Authorization"))
	if identifier == "" {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	username := r.URL.Query().Get("username")

	myId, err := strconv.Atoi(identifier)
	if err != nil {
		ctx.Logger.WithError(err).Error("error get my id")
		w.WriteHeader(http.StatusInternalServerError)
	}
	users, err := rt.db.SearchUser(myId, username)

	if err != nil {
		ctx.Logger.WithError(err).Error("error database")
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(users)

}

// Function that returns the information of the selected user
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

	// check if the user is banned
	if rt.db.IsBanned(myId, uid) || rt.db.IsBanned(uid, myId) {
		w.WriteHeader(http.StatusPartialContent)
		return
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
