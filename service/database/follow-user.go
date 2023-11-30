package database

import (
	"database/sql"
	"errors"
	"log"
)

func (db *appdbimpl) FollowUser(myId int, user string) error {
	//CHECK ROW ESISTENTE
	uid, err := db.GetId(user)
	if err != nil {
		return err
	}
	exist, _ := db.IsBanned(myId, uid)
	if exist == true {
		return ErrProfileDoesNotExist
	}

	row := db.c.QueryRow(`SELECT * FROM Follow WHERE user=? AND followedBy=?`, uid, myId)
	var id int
	exists := row.Scan(&id)
	if errors.Is(exists, sql.ErrNoRows) {
		//insert row
		_, err := db.c.Exec(`INSERT INTO Follow (user, followedBy) VALUES (?,?)`, uid, myId)
		if err != nil {
			return err
		}
		//
		row.Scan(&id)
		log.Printf("User followed: %s\n", user)
	} else {
		return ErrAlreadyFollowed
	}
	return nil
}
