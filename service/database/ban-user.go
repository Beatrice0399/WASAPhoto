package database

func (db *appdbimpl) BanUser(myId int, idProfile int) error {
	row := db.c.QueryRow(`SELECT * FROM Ban WHERE banned=? AND whoBan=?`, idProfile, myId)
	if row != nil {
		return ErrAlreadyBanned
	}

	errUn := db.UnfollowUser(myId, idProfile)
	if errUn != nil {
		return errUn
	}

	errUn = db.UnfollowUser(myId, idProfile)
	if errUn != nil {
		return errUn
	}

	_, err := db.c.Exec(`INSERT INTO Ban (banned, whoBan) VALUES (?,?)`, idProfile, myId)
	if err != nil {
		return err
	}

	return nil
}
