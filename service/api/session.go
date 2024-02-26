package api

import (
	"encoding/json"
	"net/http"

	"github.com/Beatrice0399/WASAPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("Content-Type", "application/json")

	var user User
	err := json.NewDecoder(r.Body).Decode(&user.Username)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if !validStringUsername(user.Username) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user.Uid, err = rt.db.DoLogin(user.Username)
	if err != nil {
		rt.responsError(http.StatusBadRequest, err.Error(), w)
		return
	}
	rt.responseJson(user.Uid, w)
}
