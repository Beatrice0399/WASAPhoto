package database

import (
	"database/sql"
	"log"
)

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
		//log.Printf("id: %d, name: %s\n", id, name)
		u, _ := db.GetUserProfile(id, 0)
		profiles = append(profiles, u)
		//log.Printf("id: %d, name: %s, follower: %d, following: %d, pho: %d\n", u.ID, u.Name, u.Follower, u.Following, &u.NumberPhotos)
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

func (db *appdbimpl) GetTableBan() (*sql.Rows, error) {
	rows, err := db.c.Query("SELECT * FROM Ban")
	if err != nil {
		return rows, err
	}
	return rows, nil
}

func (db *appdbimpl) GetBanned(myId int) ([]User, error) {
	var users []User
	rows, err := db.c.Query(`SELECT u.*
							FROM User u JOIN Ban b ON b.whoBan=u.id WHERE b.whoBan=?`, myId)
	if err != nil {
		log.Print("ERRORE GetBanned")
		return users, err
	}
	for rows.Next() {
		var u User
		_ = rows.Scan(&u.Uid, &u.Username)
		log.Printf("id: %d, name: %s\n", u.Uid, u.Username)
		users = append(users, u)

	}

	return users, nil
}

func (db *appdbimpl) GetTableComment() (*sql.Rows, error) {
	rows, err := db.c.Query("SELECT * FROM Comment")
	if err != nil {
		return rows, err
	}
	return rows, nil
}

func (db *appdbimpl) GetTableLikes() (*sql.Rows, error) {
	rows, err := db.c.Query("SELECT * FROM Likes")
	if err != nil {
		return rows, err
	}
	return rows, nil
}
