package database

// Database function that allows an user (fid) to follow another one (uid)
func (db *appdbimpl) FollowUser(fid int, uid int) error {
	// Non posso seguire chi mi ha bannato
	if db.IsBanned(fid, uid) {
		return ErrFollowUser
	}

	_ = db.UnbanUser(fid, uid)
	// insert row
	_, err := db.c.Exec(`INSERT INTO Follow (user, followedBy) VALUES (?,?)`, uid, fid)
	if err != nil {
		return err
	}
	return nil
}

// Database function that allows an user (fid) to unfollow another one (uid)
func (db *appdbimpl) UnfollowUser(fid int, uid int) error {
	res, err := db.c.Exec(`DELETE FROM Follow WHERE user=? AND followedBy=?`, uid, fid)
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

// Database function that returns the list of users' followers
func (db *appdbimpl) GetFollower(id int) ([]User, error) {
	rows, err := db.c.Query(`SELECT f.followedBy
							FROM Follow f
							WHERE f.user=?`, id)
	if err != nil {
		return nil, err
	}

	var users []User
	defer rows.Close()
	for rows.Next() {
		var pid int
		err = rows.Scan(&pid)
		if err != nil {
			return users, err
		}
		row := db.c.QueryRow(`SELECT * FROM User WHERE id=?`, pid)
		var u User
		err = row.Scan(&u.Uid, &u.Username)
		if err != nil {
			return users, err
		}
		// log.Printf("Function GetFollower. id: %d, name: %s\n", u.Uid, u.Username)
		users = append(users, u)
	}
	if rows.Err() != nil {
		return nil, err
	}
	return users, err
}

// Database function that returns the list of the users that the given id followed
func (db *appdbimpl) GetFollowing(followedBy int) ([]User, error) {
	rows, err := db.c.Query(`SELECT u.* FROM User u
							JOIN Follow f ON u.id=f.user WHERE f.followedBy=?`, followedBy)
	if err != nil {
		return nil, err
	}

	var users []User
	defer rows.Close()
	for rows.Next() {
		var u User
		err = rows.Scan(&u.Uid, &u.Username)
		if err != nil {
			return users, err
		}
		users = append(users, u)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return users, err
}
