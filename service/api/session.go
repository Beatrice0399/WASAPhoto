package api

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/Beatrice0399/WASAPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("Content-Type", "application/json")

	var user User
	var username Username
	err := json.NewDecoder(r.Body).Decode(&username)
	if err != nil {
		log.Println("errore qui: ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if !validStringUsername(username.Username) {
		log.Println("errore 2 qui: ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user.Uid, err = rt.db.DoLogin(username.Username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("Error login")
		return
	}

	err = createUserFolder(user.Uid, ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("can't create user's photo folder")
		return
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("can't create response json")
	}
}

func createUserFolder(identifier int, ctx reqcontext.RequestContext) error {
	path := filepath.Join(photoFolder, strconv.Itoa(identifier))

	err := os.MkdirAll(filepath.Join(path, "photos"), os.ModePerm)

	if err != nil {
		ctx.Logger.WithError(err).Error("error creating directories for user")
		return err
	}
	return nil
}
