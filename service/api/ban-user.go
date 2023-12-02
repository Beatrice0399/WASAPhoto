package api

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) banUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var username string
	username = ps.ByName("pid")
	var myId_string string
	myId_string = r.URL.Query().Get("myid")
	myId, _ := strconv.Atoi(myId_string)
	_ = rt.db.BanUser(myId, username)
}
