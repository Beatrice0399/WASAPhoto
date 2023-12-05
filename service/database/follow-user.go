package database

func (db *appdbimpl) FollowUser(myId int, user string) error {
	uid, err := db.GetId(user)
	if err != nil {
		return err
	}
	//Non posso seguire chi mi ha bannato
	exist, _ := db.IsBanned(uid, myId)
	if exist == true {
		return ErrFollowUser
	}

	_ = db.UnbanUser(myId, user)
	//insert row
	_, err = db.c.Exec(`INSERT INTO Follow (user, followedBy) VALUES (?,?)`, uid, myId)
	if err != nil {
		return err
	}
	return nil
}
