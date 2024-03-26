package database

func (db *appdbimpl) BanUser(myId int, bid int) error {

	err := db.UnfollowUser(myId, bid)
	/*
		if err != nil {
			return err
		}
	*/
	err = db.removeAllComments(myId, bid)
	/*
		if err != nil {
			return err
		}
	*/
	err = db.removeAllComments(bid, myId)
	/*
		if err != nil {
			return err
		}
	*/
	err = db.removeAllLikes(myId, bid)
	/*
		if err != nil {
			return err
		}
	*/

	err = db.removeAllLikes(bid, myId)
	/*
		if err != nil {
			return err
		}
	*/

	_ = db.UnfollowUser(bid, myId)

	_, err = db.c.Exec(`INSERT INTO Ban (banned, whoBan) VALUES (?,?)`, bid, myId)
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
