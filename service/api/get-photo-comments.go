package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getPhotoComments(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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
