package api

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var username string
	username = ps.ByName("pid")
	uid, _ := rt.db.GetId(username)
	profile, _ := rt.db.GetUserProfile(uid)

	str := fmt.Sprintf("Uid: %d, Name: %s, Follower: %d, Following: %d, Photo: NO\n", profile.ID, profile.Name, profile.Follower, profile.Following)
	w.Write([]byte(str))
}
