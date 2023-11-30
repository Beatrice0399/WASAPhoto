package database

func (db *appdbimpl) GetFollower(id int) ([]User, error) {
	rows, err := db.c.Query(`SELECT f.followedBy
							FROM Follow f
							WHERE f.user=?`, id)
	if err != nil {
		return nil, err
	}

	var users []User
	for rows.Next() {
		var pid int
		err = rows.Scan(&pid)
		if err != nil {
			return nil, err
		}
		row := db.c.QueryRow(`SELECT * FROM User WHERE id=?`, pid)
		var u User
		err = row.Scan(&u.ID, &u.Name)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, err
}
