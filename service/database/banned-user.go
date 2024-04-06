package database

func (db *appdbimpl) BannedUser(myId int) ([]User, error) {
	var banned []User
	rows, err := db.c.Query(`SELECT u FROM Ban b, User u WHERE whoBan=?`, myId)
	if err != nil {
		return banned, err
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		err = rows.Scan(&user.Uid, &user.Username)
		if err != nil {
			return banned, err
		}
		banned = append(banned, user)
	}
	if err = rows.Err(); err != nil {
		return banned, err
	}
	if rows.Err() != nil {
		return nil, err
	}

	return banned, nil
}
