package database

// Database function that searches users with the given username
func (db *appdbimpl) SearchUser(myId int, username string) ([]User, error) {
	rows, err := db.c.Query(`SELECT *
							FROM User
							WHERE username LIKE '%' || ? || '%' AND id NOT IN (SELECT whoBan FROM ban b JOIN user u ON u.id=b.whoBan WHERE (b.banned=? AND b.whoBan=u.id))`, username, myId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []User
	for rows.Next() {
		var u User
		err = rows.Scan(&u.Uid, &u.Username)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	if rows.Err() != nil {
		return nil, err
	}
	return users, err
}

// Database function that modifies the username with the newone given
func (db *appdbimpl) SetMyUsername(id int, name string) error {
	res, err := db.c.Exec(`UPDATE User SET username=? WHERE id=?`, name, id)
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

// Database function that return the user's profile
func (db *appdbimpl) GetUserProfile(id int, myId int) (Profile, error) {
	var profile Profile
	if db.IsBanned(id, myId) {
		return profile, ErrUserBannedYou
	}
	if db.IsBanned(myId, id) {
		return profile, ErrUserBanned
	}
	row := db.c.QueryRow(`SELECT * FROM User WHERE id=?`, id)

	err := row.Scan(&profile.ID, &profile.Name)
	if err != nil {
		return profile, ErrProfileDoesNotExist
	}

	follower, err := db.GetFollower(id)
	if err != nil {
		return profile, err
	}
	profile.Followers = follower

	following, err := db.GetFollowing(id)
	if err != nil {
		return profile, err
	}
	profile.Following = following

	// profile.NumberPhotos, _ = db.GetNumberPhotoUser(id)
	profile.Photos, _ = db.GetPhotoUser(id)

	return profile, err
}
