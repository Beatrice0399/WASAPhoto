package database

import "log"

func (db *appdbimpl) UnbanUser(myId int, user string) error {
	idProfile, _ := db.GetId(user)
	res, err := db.c.Exec(`DELETE FROM Ban WHERE banned=? AND whoBan=?`, idProfile, myId)
	if err != nil {
		log.Println("Errore durante l'unban:", err)
		return err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return err
	} else if affected == 0 {
		return ErrProfileDoesNotExist
	}
	return nil
}
