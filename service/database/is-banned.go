package database

func (db *appdbimpl) IsBanned(myId int, idProfile int) bool {
	row := db.c.QueryRow(`SELECT * FROM Banned WHERE banned=? AND whoBan=?`, idProfile, myId)
	var text string
	err := row.Scan(&text)
	if err != nil {
		return false
	} else {
		return true
	}
}
