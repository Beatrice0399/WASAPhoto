package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) setMyUsername(w http.ResponseWriter, r *http.Request, pd httprouter.Params) {
	var username string
	err := json.NewDecoder(r.Body).Decode(&username)
	if err != nil {
		w.Header().Set("Content-type", "application/json")
		w.Write([]byte("Invalid username login"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var myId_string string
	myId_string = r.URL.Query().Get("myid")
	myId, _ := strconv.Atoi(myId_string)
	log.Printf("SetUsername. id: %d, newName: %s", myId, username)
	name, err := rt.db.SetMyUsername(myId, username)
	w.Header().Set("Content-type", "application/json")
	id_byte, err := json.Marshal(name)
	w.Write(id_byte)

}
