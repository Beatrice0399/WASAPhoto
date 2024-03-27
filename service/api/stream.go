package api

import (
	"encoding/json"
	"net/http"

	"github.com/Beatrice0399/WASAPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getMyStream(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("Content-Type", "application/json")

	// A user can only see his/her home
	valid := validateRequestingUser(ps.ByName("uid"), extractBearer(r.Header.Get("Authorization")))
	if valid != 0 {
		w.WriteHeader(valid)
		return
	}

	myid, err := rt.get_myid_path(ps)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	stream, err := rt.db.GetMyStream(myid)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(stream)
}
