package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/julienschmidt/httprouter"
)

func validStringUsername(identifier string) bool {
	var trimmedId = strings.TrimSpace(identifier)
	return len(identifier) >= 3 && len(identifier) <= 16 && trimmedId != "" && !strings.ContainsAny(trimmedId, "?_")
}

func (rt *_router) getProfiles(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	rows, err := rt.db.GetAllProfiles()
	if err != nil {
		return
	}
	w.Header().Set("Content-type", "application/json")
	for _, value := range rows {
		//str := fmt.Sprintf("id: %d, Name: %s, Follower: %d, Following: %d, Photo: %d\n", value.ID, value.Name, value.Followers, value.Following, value.NumberPhotos)
		//w.Write([]byte(str))
		_ = json.NewEncoder(w).Encode(value)

	}
}

func extractBearer(authorization string) string {
	var tokens = strings.Split(authorization, " ")
	if len(tokens) == 2 {
		return strings.Trim(tokens[1], " ")
	}
	return ""
}

func isNotLogged(auth string) bool {

	return auth == ""
}

func validateRequestingUser(identifier string, bearerToken string) int {

	// If the requesting user has an invalid token then respond with a fobidden status
	if isNotLogged(bearerToken) {
		return http.StatusForbidden
	}

	//  If the requesting user's id is different than the one in the path then respond with a unathorized status.

	if identifier != bearerToken {
		return http.StatusUnauthorized
	}
	return 0
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
		log.Printf("id: %d, name: %s\n", user.Uid, user.Username)
		str := fmt.Sprintf("id: %d, name: %s\n", user.Uid, user.Username)
		w.Write([]byte(str))
	}
}

func (rt *_router) getTableComment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	rows, err := rt.db.GetTableComment()
	if err != nil {
		rt.baseLogger.Errorln(err)
	}
	var id int
	var user int
	var photo int
	var txt string
	var date string
	var boolean int
	w.Header().Set("Content-type", "application/json")
	for exist := rows.Next(); exist == true; exist = rows.Next() {
		err = rows.Scan(&id, &user, &photo, &txt, &date, &boolean)
		if err != nil {
			rt.baseLogger.Errorln(err)
		}
		str := fmt.Sprintf("cid: %d, user: %d, photo: %d, txt: %s, date: %s, bool: %d \n", id, user, photo, date, txt, boolean)
		w.Write([]byte(str))
	}
}

func (rt *_router) getTableLikes(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	rows, err := rt.db.GetTableLikes()
	if err != nil {
		rt.baseLogger.Errorln(err)
	}
	var id int
	var phid int
	var uid int
	for exist := rows.Next(); exist == true; exist = rows.Next() {
		err = rows.Scan(&id, &phid, &uid)
		if err != nil {
			rt.baseLogger.Errorln(err)
		}
		str := fmt.Sprintf("lid: %d, photo: %d, user: %d\n", id, phid, uid)
		w.Write([]byte(str))
	}
}
