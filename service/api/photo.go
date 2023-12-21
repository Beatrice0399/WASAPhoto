package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/Beatrice0399/WASAPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) deletePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	phid, err := rt.getPhid(ps)
	if err != nil {
		rt.responsError(http.StatusBadRequest, err.Error(), w)
		return
	}

	myid, err := rt.get_myid_path(ps)
	if err != nil {
		rt.responsError(http.StatusBadRequest, err.Error(), w)
		return
	}

	err = rt.db.DeletePhoto(phid, myid)
	if err != nil {
		rt.responsError(http.StatusBadRequest, err.Error(), w)
		return
	}
	str := "Photo deleted"
	rt.responseJson(str, w)
}

func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var img string
	err := json.NewDecoder(r.Body).Decode(&img)
	if err != nil {
		rt.responsError(http.StatusBadRequest, err.Error(), w)
		return
	}
	myId, err := rt.get_myid_path(ps)
	if err != nil {
		rt.responsError(http.StatusBadRequest, err.Error(), w)
		return
	}

	p, err := rt.db.UploadPhoto(myId, []byte(img))
	if err != nil {
		rt.responsError(http.StatusBadRequest, err.Error(), w)
		return
	}
	rt.responseJson(p, w)
}

func (rt *_router) getPhotoComments(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var phid_string string
	phid_string = ps.ByName("phid")
	phid, _ := strconv.Atoi(phid_string)

	comments, err := rt.db.GetPhotoComments(phid)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(comments)
}
