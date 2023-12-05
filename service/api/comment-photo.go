package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) commentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var phid_string string
	phid_string = ps.ByName("phid")
	phid, _ := strconv.Atoi(phid_string)

	var myId_string string
	myId_string = r.URL.Query().Get("myid")
	myId, _ := strconv.Atoi(myId_string)

	var msg string
	err := json.NewDecoder(r.Body).Decode(&msg)

	if err != nil {
		w.Header().Set("Content-type", "application/json")
		w.Write([]byte("Invalid username login"))
		log.Print("error deconding body message")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cid, _ := rt.db.CommentPhoto(myId, phid, msg)
	w.Header().Set("Content-type", "application/json")
	id_byte, err := json.Marshal(cid)
	w.Write(id_byte)
	log.Printf("commento: %s, id: %d", msg, cid)
}
