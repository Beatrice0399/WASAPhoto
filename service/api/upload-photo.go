package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var img string
	err := json.NewDecoder(r.Body).Decode(&img)
	if err != nil {
		w.Header().Set("Content-type", "application/json")
		w.Write([]byte("Invalid username login"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var myId_string string
	myId_string = r.URL.Query().Get("myid")
	myId, _ := strconv.Atoi(myId_string)
	_, _ = rt.db.UploadPhoto(myId, []byte(img))
}
