package database

func (db *appdbimpl) GetFollowing(followedBy int) ([]User, error) {
	rows, err := db.c.Query(`SELECT u.* FROM User u
							JOIN Follow f ON u.id=f.followedBy WHERE f.followedBy=?`, followedBy)
	if err != nil {
		return nil, err
	}

	var users []User
	for rows.Next() {
		var u User
		err = rows.Scan(&u.ID, &u.Name)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, err
}
