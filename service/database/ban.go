package database

func (db *appdbimpl) BanUser(myId int, bid int) error {
	_ = db.UnfollowUser(myId, bid)
	_ = db.removeAllComments(myId, bid)
	_ = db.removeAllComments(bid, myId)
	_ = db.removeAllLikes(myId, bid)
	_ = db.removeAllLikes(bid, myId)
	_ = db.UnfollowUser(bid, myId)

	_, err := db.c.Exec(`INSERT INTO Ban (banned, whoBan) VALUES (?,?)`, bid, myId)
	if err != nil {
		return err
	}

	return nil
}

func (db *appdbimpl) UnbanUser(myId int, bid int) error {
	res, err := db.c.Exec(`DELETE FROM Ban WHERE banned=? AND whoBan=?`, bid, myId)
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
