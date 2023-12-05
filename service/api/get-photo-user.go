package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getPhotoUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var myId_string string
	myId_string = r.URL.Query().Get("myid")
	myId, _ := strconv.Atoi(myId_string)

	rows, err := rt.db.GetPhotoUser(myId)
	if err != nil {
		rt.baseLogger.Errorln(err)
	}
	//var p database.Photo
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(rows)
	/*
		for exist := rows.Next(); exist == true; exist = rows.Next() {
			err = rows.Scan(&p.ID, &p.User, &p.Image, &p.Date)
			if err != nil {
				rt.baseLogger.Errorln(err)
			}
			likes, _ := rt.db.GetLikesPhoto(p.ID)
			str := fmt.Sprintf("pid: %d, User: %d, img: %s, date: %s, likes: %d\n", p.ID, p.User, p.Image, p.Date, likes)
			w.Write([]byte(str))
		}
	*/
}
