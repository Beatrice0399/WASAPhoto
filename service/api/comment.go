package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Beatrice0399/WASAPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

// function that adds a comment to a photo and return the created comment
func (rt *_router) commentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("Content-Type", "application/json")

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

	phid, err := rt.getPhid(ps)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var msg CommentText
	err = json.NewDecoder(r.Body).Decode(&msg)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		ctx.Logger.WithError(err).Error("error decoding request body json")
		return
	}
	if len(msg.Comment) > 400 {
		w.WriteHeader(http.StatusBadRequest)
		ctx.Logger.WithError(err).Error("error comment longer than 400 characters")
		return
	}

	comment, err := rt.db.CommentPhoto(myID, phid, msg.Comment)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("error insert comment within databse")
		return
	}
	w.WriteHeader(http.StatusCreated)

	err = json.NewEncoder(w).Encode(comment)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("error convert id photo")
		return
	}

}

// function that allows a user to delete a comment
func (rt *_router) uncommentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	token := extractBearer(r.Header.Get("Authorization"))
	// Check if the user isn't logged
	if isNotLogged(token) {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	cid, err := rt.getCid(ps)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	phid, err := rt.getPhid(ps)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	myID, err := strconv.Atoi(token)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	err = rt.db.UncommentPhoto(cid, phid, myID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
