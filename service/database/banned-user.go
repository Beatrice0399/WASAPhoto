package database

func (db *appdbimpl) BannedUser(myId int) ([]User, error) {
	var ret []User
	rows, err := db.c.Query(`SELECT u FROM Ban b, User u WHERE whoBan=?`, myId)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	for rows.Next() {
		var user User
		err = rows.Scan(&user.Uid, &user.Username)
		if err != nil {
			return nil, err
		}
		ret = append(ret, user)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return ret, nil
}
