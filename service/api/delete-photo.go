package api

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) deletePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var phid_string string
	phid_string = ps.ByName("phid")
	phid, _ := strconv.Atoi(phid_string)

	_ = rt.db.DeletePhoto(phid)
}
