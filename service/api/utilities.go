package api

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getProfiles(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	rows, err := rt.db.GetAllProfiles()
	if err != nil {
		return
	}
	w.Header().Set("Content-type", "application/json")
	for _, value := range rows {
		str := fmt.Sprintf("id: %d, Name: %s, Follower: %d, Following: %d, Photo: %d\n", value.ID, value.Name, value.Follower, value.Following, value.NumberPhotos)
		w.Write([]byte(str))
	}
}

func (rt *_router) getUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	rows, err := rt.db.GetAllUsers()
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

func (rt *_router) getFollows(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	rows, err := rt.db.GetTableFollow()
	if err != nil {
		rt.baseLogger.Errorln(err)
	}
	var id int
	var name int
	var whoFol int
	w.Header().Set("Content-type", "application/json")
	for exist := rows.Next(); exist == true; exist = rows.Next() {
		err = rows.Scan(&id, &name, &whoFol)
		if err != nil {
			rt.baseLogger.Errorln(err)
		}
		str := fmt.Sprintf("Uid: %d, Follow: %d, WhoFollow: %d\n", id, name, whoFol)
		w.Write([]byte(str))
	}
}

func (rt *_router) getTableBan(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	rows, err := rt.db.GetTableBan()
	if err != nil {
		rt.baseLogger.Errorln(err)
	}
	var id int
	var name int
	var whoBan int
	w.Header().Set("Content-type", "application/json")
	for exist := rows.Next(); exist == true; exist = rows.Next() {
		err = rows.Scan(&id, &name, &whoBan)
		if err != nil {
			rt.baseLogger.Errorln(err)
		}
		str := fmt.Sprintf("Uid: %d, Banned: %d, WhoBan: %d\n", id, name, whoBan)
		w.Write([]byte(str))
	}
}

func (rt *_router) getBanned(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var myId_string string
	myId_string = r.URL.Query().Get("myid")
	myId, _ := strconv.Atoi(myId_string)
	users, _ := rt.db.GetBanned(myId)
	for _, user := range users {
		log.Printf("id: %d, name: %s\n", user.ID, user.Name)
		str := fmt.Sprintf("id: %d, name: %s\n", user.ID, user.Name)
		w.Write([]byte(str))
	}
}
