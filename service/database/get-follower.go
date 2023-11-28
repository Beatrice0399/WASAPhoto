package database

func (db *appdbimpl) GetFollower(id uint64) ([]User, error) {
	rows, err := db.c.Query(`SELECT followed FROM Follow WHERE user=?`, id)
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
