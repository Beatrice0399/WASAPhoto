package database

func (db *appdbimpl) IsBanned(myId int, idProfile int) (bool, error) {
	row := db.c.QueryRow(`SELECT * FROM Banned WHERE banned=? AND whoBan=?`, idProfile, myId)
	text := ""
	err := row.Scan(&text)
	if err != nil {
		return false, nil
	} else {
		return true, nil
	}
}
