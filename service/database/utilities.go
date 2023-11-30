package database

import "database/sql"

func (db *appdbimpl) GetAllProfiles() ([]Profile, error) {
	var profiles []Profile
	rows, err := db.c.Query(`SELECT * FROM User`)
	if err != nil {
		return profiles, nil
	}
	var id int
	var name string
	for rows.Next() {
		err = rows.Scan(&id, &name)
		u, _ := db.GetUserProfile(id)
		profiles = append(profiles, u)
		/*
			err = rows.Scan(&u.ID, &u.Name, &u.Follower, &u.Following, &u.NumberPhotos)
			if err != nil {
				return profiles, err
			}
			id, _ := db.GetId(u.Name)
			follower, _ := db.GetFollower(id)
			following, _ := db.GetFollowing(id)
			photos, _ := db.GetPhotoUser(id)
			u.Follower = len(follower)
			u.Following = len(following)
			u.NumberPhotos = len(photos)
			u.Photos = photos

			profiles = append(profiles, u)
		*/
	}

	return profiles, err
}

func (db *appdbimpl) GetAllUsers() (*sql.Rows, error) {
	rows, err := db.c.Query("SELECT * FROM User")
	if err != nil {
		return rows, err
	}
	return rows, nil
}

func (db *appdbimpl) GetTableFollow() (*sql.Rows, error) {
	rows, err := db.c.Query("SELECT * FROM Follow")
	if err != nil {
		return rows, err
	}
	return rows, nil
}
