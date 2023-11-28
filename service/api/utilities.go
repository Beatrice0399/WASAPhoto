package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Beatrice0399/WASAPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) GetAllUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	users, err := rt.db.GetAllUsers()
	if err != nil {
		return
	}
	w.Header().Set("Content-type", "application/json")
	users_byte, err := json.MarshalIndent(users, "", "\t")
	w.Write(users_byte)
	return
}

func (rt *_router) getProfiles(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	rows, err := rt.db.GetAllProfiles()
	if err != nil {
		rt.baseLogger.Errorln(err)
	}
	var uid int
	var name string
	w.Header().Set("Content-type", "application/json")
	for exist := rows.Next(); exist == true; exist = rows.Next() {
		err = rows.Scan(&uid, &name)
		if err != nil {
			rt.baseLogger.Errorln(err)
		}
		str := fmt.Sprintf("Uid: %d, Name: %s\n", uid, name)
		w.Write([]byte(str))
	}
}
