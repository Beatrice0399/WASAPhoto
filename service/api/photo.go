package api

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/Beatrice0399/WASAPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

// Function to upload a photo
func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("Content-Type", "application/json")

	// Check the user's identity for the operation
	valid := validateRequestingUser(ps.ByName("uid"), extractBearer(r.Header.Get("Authorization")))
	if valid != 0 {
		w.WriteHeader(valid)
		return
	}

	data, err := io.ReadAll(r.Body)
	if err != nil {
		ctx.Logger.WithError(err).Error("photo-upload: error reading body content")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// After reading the body we won't be able to read it again. We'll reassign a "fresh" io.ReadCloser to the body
	r.Body = io.NopCloser(bytes.NewBuffer(data))

	err = checkFormatPhoto(r.Body, io.NopCloser(bytes.NewBuffer(data)))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		ctx.Logger.WithError(err).Error("photo-upload: body contains file that is neither jpeg or png")
		_ = json.NewEncoder(w).Encode(ErrMsgJSON{Message: FORMAT_ERROR_IMG})
		return
	}

	r.Body = io.NopCloser(bytes.NewBuffer(data))

	uid, err := rt.get_uid_path(ps)
	if err != nil {
		ctx.Logger.WithError(err).Error("error get user id from path")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	phid, err := rt.db.NewPhoto(uid)
	if err != nil {
		ctx.Logger.WithError(err).Error("error creating new photo within database")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	path, err := getPhotoFolder(ps.ByName("uid"))
	if err != nil {
		ctx.Logger.WithError(err).Error("error getting user's photo folder")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	completePath := filepath.Join(path, strconv.Itoa(phid))
	result, err := os.Create(completePath)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("error creating local photo file")
		return
	}
	_, err = io.Copy(result, r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("error copying body content into file photo")
		return
	}

	result.Close()
	username, err := rt.db.GetNameById(uid)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("error update photo's path within the database")
		return
	}
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(Photo{
		Phid:     phid,
		User:     uid,
		Username: username,
		Path:     completePath,
		Date:     time.Now().UTC(),
	})

}

// Functions that deletes a photo and its comments and likes
func (rt *_router) deletePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	token := extractBearer(r.Header.Get("Authorization"))
	// check the user's identity for the operation
	valid := validateRequestingUser(ps.ByName("uid"), token)
	if valid != 0 {
		w.WriteHeader(valid)
		return
	}

	myid, err := rt.get_uid_path(ps)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	phid, err := rt.getPhid(ps)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = rt.db.DeletePhoto(phid, myid)
	if err != nil {
		ctx.Logger.WithError(err).Error("error removing phot from database")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	path, err := getPhotoFolder(token)
	if err != nil {
		ctx.Logger.WithError(err).Error("photo-delete/getUserPhotoFolder: error with directories")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Remove the file from the user's photos folder
	err = os.Remove(filepath.Join(path, ps.ByName("phid")))
	if err != nil {
		ctx.Logger.WithError(err).Error("photo to be removed not found")
	}

	w.WriteHeader(http.StatusNoContent)

}

// Returns the path of the photo
func (rt *_router) getPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	http.ServeFile(w, r,
		filepath.Join(photoFolder, ps.ByName("uid"), "photos", ps.ByName("phid")))
}
