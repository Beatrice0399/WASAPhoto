package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getMyStream(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var myId_string string
	myId_string = r.URL.Query().Get("myid")
	myId, _ := strconv.Atoi(myId_string)

	stream, _ := rt.db.GetMyStream(myId)
	w.Header().Set("Content-type", "application/json")

	json.NewEncoder(w).Encode(stream)

}
