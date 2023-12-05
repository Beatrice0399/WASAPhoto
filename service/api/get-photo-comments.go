package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getPhotoComments(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var phid_string string
	phid_string = ps.ByName("phid")
	phid, _ := strconv.Atoi(phid_string)

	comments, err := rt.db.GetPhotoComments(phid)
	if err != nil {
		log.Println(err)
	}

	for _, c := range comments {
		//str := fmt.Sprintf("cid: %d, User: %s, text: %s, date: %s\n", c.ID, c.User, c.Text, c.Date)
		//w.Write([]byte(str))
		json.NewEncoder(w).Encode(c)
	}
}
