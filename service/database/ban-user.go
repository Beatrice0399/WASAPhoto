package database

import "log"

func (db *appdbimpl) BanUser(myId int, username string) error {
	idProfile, err := db.GetId(username)
	if err != nil {
		return err
	}
	/*
		row := db.c.QueryRow(`SELECT * FROM Ban WHERE banned=? AND whoBan=?`, idProfile, myId)
		if row != nil {
			log.Print("Errore BanUser in queryRow")
			return ErrAlreadyBanned
		}
	*/
	_ = db.UnfollowUser(myId, username)
	/*
		if errUn != nil {
			log.Print("Errore BanUser in unfollofUser1")
			return errUn
		}
	*/
	name, _ := db.GetNameById(myId)
	_ = db.UnfollowUser(idProfile, name)
	/*
		if errUn != nil {
			log.Print("Errore BanUser in unfollofUser2")
			return errUn
		}
	*/
	_, err = db.c.Exec(`INSERT INTO Ban (banned, whoBan) VALUES (?,?)`, idProfile, myId)
	if err != nil {
		log.Println("Errore insert:", err)
		return err
	}

	return nil
}
