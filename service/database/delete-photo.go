package database

import "log"

func (db *appdbimpl) DeletePhoto(pid int) error {
	res, err := db.c.Exec(`DELETE FROM Photo WHERE id=?`, pid)
	if err != nil {
		log.Println("ERR DeletePhoto: ", err)
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
