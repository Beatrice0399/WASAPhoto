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
	var new_username Username
	err := json.NewDecoder(r.Body).Decode(&new_username)
	log.Println(new_username.Username)
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
	uid := ps.ByName("uid")

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
	err = rt.db.SetMyUsername(myId, new_username.Username)
	if err != nil {
		ctx.Logger.WithError(err).Error("error executing query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

	// newURL := "/users/" + new_username
	// http.Redirect(w, r, newURL, http.StatusFound)
}

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

// ritorna le informazioni del profilo dell'utente cercato
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

	if rt.db.IsBanned(myId, uid) || rt.db.IsBanned(uid, myId) {
		ctx.Logger.WithError(err).Error("Profile unavailable")
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
