package database

func (db *appdbimpl) BanUser(myId int, username string) error {
	idProfile, err := db.GetId(username)
	if err != nil {
		return err
	}

	_ = db.UnfollowUser(myId, username)

	name, err := db.GetNameById(myId)
	if err != nil {
		return err
	}
	_ = db.UnfollowUser(idProfile, name)

	_, err = db.c.Exec(`INSERT INTO Ban (banned, whoBan) VALUES (?,?)`, idProfile, myId)
	if err != nil {
		return err
	}

	return nil
}

func (db *appdbimpl) UnbanUser(myId int, user string) error {
	idProfile, _ := db.GetId(user)
	res, err := db.c.Exec(`DELETE FROM Ban WHERE banned=? AND whoBan=?`, idProfile, myId)
	if err != nil {
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
