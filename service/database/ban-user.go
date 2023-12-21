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
