package api

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) followUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var username string
	username = ps.ByName("pid")
	/*
		id, err := rt.db.GetId(username)
		if err != nil {
			w.Header().Set("Content-type", "application/json")
			w.Write([]byte("Invalid username login"))
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	*/
	var myId_string string
	myId_string = r.URL.Query().Get("myid")
	myId, _ := strconv.Atoi(myId_string)
	_ = rt.db.FollowUser(myId, username)

}
