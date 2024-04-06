package api

import (
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) get_uid_path(ps httprouter.Params) (int, error) {
	username := ps.ByName("uid")
	myid, err := strconv.Atoi(username)
	if err != nil {
		return -1, err
	}
	return myid, nil
}

/*
	func (rt *_router) get_uid_query(r *http.Request) (int, error) {
		myid, err := strconv.Atoi(r.URL.Query().Get("uid"))
		if err != nil {
			return 0, err
		}
		return myid, nil
	}
*/

func (rt *_router) getPhid(ps httprouter.Params) (int, error) {
	string_pid := ps.ByName("phid")
	pid, err := strconv.Atoi(string_pid)
	if err != nil {
		return 0, err
	}
	return pid, nil
}

func (rt *_router) getCid(ps httprouter.Params) (int, error) {
	cid_string := ps.ByName("cid")
	cid, err := strconv.Atoi(cid_string)
	if err != nil {
		return 0, err
	}
	return cid, nil
}

func (rt *_router) get_like_id(ps httprouter.Params) (int, error) {
	lid_string := ps.ByName("lid")
	lid, err := strconv.Atoi(lid_string)
	if err != nil {
		return 0, err
	}
	return lid, nil
}
