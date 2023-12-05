package api

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) uncommentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var cid_string string
	cid_string = ps.ByName("cid")
	cid, _ := strconv.Atoi(cid_string)

	var phid_string string
	phid_string = ps.ByName("phid")
	phid, _ := strconv.Atoi(phid_string)

	var myId_string string
	myId_string = r.URL.Query().Get("myid")
	myId, _ := strconv.Atoi(myId_string)

	_ = rt.db.UncommentPhoto(cid, phid, myId)
}
