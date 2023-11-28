package database

import (
	"database/sql"
	"errors"
	"fmt"
)

/*
type DB struct {
	Filename string `conf:""`
}

//Open IS USED TO CREATE A DB HANDLE
func OpenDBConnection() (*sql.DB, error) {
	logger.Println("inizializing database support")
	db, err := sql.Open("sqlite3", "dataSourceName")
	if err != nil {
		logger.WithError(err).Error("error opening SQLite DB")
		return db, fmt.Errorf("opening SQLite: %w", err)
	}
	defer func() {
		logger.Debug("database stopping")
		_ = db.Close()
	}()
	return db, nil
}
*/

type AppDatabase interface {
	DoLogin(username string) (uint64, error)
	MyProfile() (Profile, error)
	SetMyUsername(id uint64, name string) error
	UploadPhoto(id uint64, img []byte) (Photo, error)
	FollowUser(myId uint64, idProfile uint64) error
	UnfollowUser(myId uint64, idProfile uint64) error
	BanUser(myId uint64, idProfile uint64) error
	BannedUser(myId uint64) ([]User, error)
	UnbanUser(myId uint64, idProfile uint64) error
	GetUserProfile(id uint64) (Profile, error)
	GetMyStream(myId uint64) ([]Photo, error)
	LikePhoto(phId uint64, uid uint64) error
	//UnlikePhoto(id uint64) error
	//CommentPhoto(id uint64, text string) (Comment, error)
	//UncommentPhoto(id uint64) error
	//DeletePhoto(id uint64) error
	IsBanned(myId uint64, idProfile uint64) (bool, error)
	GetPhotoUser(id uint64) ([]Photo, error)
	GetPhotoComments(phId uint64) ([]Comment, error)
	GetFollower(id uint64) ([]User, error)
	GetFollowing(followedBy uint64) ([]User, error)

	Ping() error

	//utilities function
	GetAllUsers() ([]User, error)
	GetAllProfiles() (*sql.Rows, error)
}

var ErrProfileDoesNotExist = errors.New("Profile doesn't exist")
var ErrAlreadyFollowed = errors.New("Profile already followed")
var ErrAlreadyBanned = errors.New("Profile already banned")
var ErrAlreadyLiked = errors.New("Already liked")
var ErrUsernameUsed = errors.New("Username already used")

type appdbimpl struct {
	c *sql.DB
}

func New(db *sql.DB) (AppDatabase, error) {
	if db == nil {
		return nil, errors.New("Database is required whem building a AppDatabase")
	}
	//controllare se va qui
	db.Exec("PRAGMA foreign_keys ON;")

	var User string
	err := db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='User';`).Scan(&User)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE IF NOT EXISTS User (id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, username TEXT NOT NULL);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure User: %w", err)
		}
	}

	var Profile string
	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='Profile';`).Scan(&Profile)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE IF NOT EXISTS Profile (id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, user INTEGER NOT NULL, follower INTEGER, following INTEGER
			nPhoto INTEGER, FOREIGN KEY (user) REFERENCES User(id));`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure Profile: %w", err)
		}
	}

	var Photo string
	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='Photo';`).Scan(&Photo)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE IF NOT EXISTS Photo (id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, user INTEGER NOT NULL, image BLOB NOT NULL,
			date TEXT NOT NULL, like INTEGER,
			FOREIGN KEY (user) REFERENCES User(id));`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure Photo: %w", err)
		}
	}

	var Comment string
	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='Comment';`).Scan(&Comment)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE IF NOT EXISTS Comment (id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, user INTEGER NOT NULL, photo INTEGER NOT NULL, string TEXT NOT NULL, 
			date TEXT NOT NULL, FOREIGN KEY (user) REFERENCES User(id));`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure Comment: %w", err)
		}
	}

	var Follow string
	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='Follow';`).Scan(&Follow)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE IF NOT EXISTS Follow (id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, user INTEGER NOT NULL, followedBy INTEGER NOT NULL,
			FOREIGN KEY (user) REFERENCES User(id), FOREIGN KEY (followedBy) REFERENCES User(id));`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure Follow: %w", err)
		}
	}

	var Ban string
	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='Ban';`).Scan(&Ban)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE IF NOT EXISTS Ban (id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, banned INTEGER NOT NULL, whoBan INTEGER NOT NULL,
			FOREIGN KEY (banned) REFERENCES User(id), FOREIGN KEY (whoBan) REFERENCES User(id));`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure Ban: %w", err)
		}
	}

	var Likes string
	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='Likes';`).Scan(&Likes)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE IF NOT EXISTS Likes (id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, phId INTEGER NOT NULL, uid INTEGER NOT NULL,
					FOREIGN KEY (phId) REFERENCES Photo(id), FOREIGN KEY (uid) REFERENCES User(id));`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure Likes: %w", err)
		}
	}

	return &appdbimpl{
		c: db,
	}, nil
}

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}
