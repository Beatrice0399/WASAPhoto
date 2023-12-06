package database

func (db *appdbimpl) GetNumberPhotoUser(myid int) (int, error) {
	res := db.c.QueryRow(`SELECT COUNT(*) FROM Photo p WHERE p.user=?`, myid)
	var likes int
	err := res.Scan(&likes)
	if err != nil {
		return -1, err
	}
	return likes, nil
}
