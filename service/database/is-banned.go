package database

// ritorna true se l'utente con idProfile è stato bannato dall'utente con myId
func (db *appdbimpl) IsBanned(bid int, uid int) bool {
	row := db.c.QueryRow(`SELECT count(*) FROM Ban WHERE banned=? AND whoBan=?`, uid, bid)
	var count int
	err := row.Scan(&count)
	if err != nil {
		return false
	}
	if count > 0 {
		return true
	} else {
		return false
	}
}
