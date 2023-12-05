package database

import "log"

func (db *appdbimpl) UncommentPhoto(cid int, phid int, uid int) error {
	res, err := db.c.Exec(`DELETE FROM Comment WHERE id=? AND user=? AND photo=?`, cid, uid, phid)
	if err != nil {
		log.Println(err)
		return err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		log.Println(err)
		return err
	} else if affected == 0 {
		log.Print("affected == 0")
		return ErrProfileDoesNotExist
	}
	return nil
}
