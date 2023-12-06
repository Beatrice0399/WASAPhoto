package api

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getMyId(r *http.Request) (int, error) {
	myid, err := strconv.Atoi(r.URL.Query().Get("myid"))
	if err != nil {
		return 0, err
	}
	return myid, nil
}

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
