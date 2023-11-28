package database

func (db *appdbimpl) FollowUser(myId uint64, idProfile uint64) error {
	//CHECK ROW ESISTENTE
	exist, _ := db.IsBanned(myId, idProfile)
	if exist == true {
		return ErrProfileDoesNotExist
	}

	row := db.c.QueryRow(`SELECT * FROM Follow WHERE user=? AND followedBy=?`, idProfile, myId)
	if row != nil {
		return ErrAlreadyFollowed
	}
	//insert row
	_, err := db.c.Exec(`INSERT INTO Follow (user, followedBy) VALUES (?,?)`, idProfile, myId)
	if err != nil {
		return err
	}

	return nil
}
