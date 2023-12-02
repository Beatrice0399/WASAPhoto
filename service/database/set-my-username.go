package database

import "log"

func (db *appdbimpl) SetMyUsername(id int, name string) (string, error) {
	res, err := db.c.Exec(`UPDATE User SET username=? WHERE id=?`, name, id)
	if err != nil {
		if err != nil {
			log.Println("Errore durante l'aggiornamento:", err)
			return name, err
		}
		return name, err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return name, err
	} else if affected == 0 {
		return name, ErrProfileDoesNotExist
	}
	return name, nil
}
