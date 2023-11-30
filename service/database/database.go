package database

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
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
	DoLogin(username string) (int, error)
	MyProfile() (Profile, error)
	SetMyUsername(id int, name string) error
	UploadPhoto(id int, img []byte) (Photo, error)
	FollowUser(myId int, username string) error
	UnfollowUser(myId int, idProfile int) error
	BanUser(myId int, idProfile int) error
	BannedUser(myId int) ([]User, error)
	UnbanUser(myId int, idProfile int) error
	GetUserProfile(id int) (Profile, error)
	GetMyStream(myId int) ([]Photo, error)
	LikePhoto(phId int, uid int) error
	//UnlikePhoto(id int) error
	//CommentPhoto(id int, text string) (Comment, error)
	//UncommentPhoto(id int) error
	//DeletePhoto(id int) error

	IsBanned(myId int, idProfile int) (bool, error)
	GetPhotoUser(id int) ([]Photo, error)
	GetPhotoComments(phId int) ([]Comment, error)
	GetFollower(id int) ([]User, error)
	GetFollowing(followedBy int) ([]User, error)
	GetId(username string) (int, error)

	Ping() error

	//utilities function
	GetAllProfiles() ([]Profile, error)
	GetAllUsers() (*sql.Rows, error)
	GetTableFollow() (*sql.Rows, error)
}

var ErrProfileDoesNotExist = errors.New("Profile doesn't exist")
var ErrAlreadyFollowed = errors.New("Profile already followed")
var ErrAlreadyBanned = errors.New("Profile already banned")
var ErrAlreadyLiked = errors.New("Already liked")
var ErrUsernameUsed = errors.New("Username already used")
var ErrWithForeignKey = errors.New("Error turning on foreign key")

type appdbimpl struct {
	c *sql.DB
}

func New(db *sql.DB) (AppDatabase, error) {
	if db == nil {
		return nil, errors.New("Database is required whem building a AppDatabase")
	}

	//DROP TABLE
	tableName := "User" // Sostituisci con il nome effettivo della tua tabella
	_, erro := db.Exec("DROP TABLE IF EXISTS " + tableName)
	if erro != nil {
		fmt.Println(erro)
		return nil, erro
	}
	tableName = "Profile" // Sostituisci con il nome effettivo della tua tabella
	_, erro = db.Exec("DROP TABLE IF EXISTS " + tableName)
	if erro != nil {
		fmt.Println(erro)
		return nil, erro
	}
	tableName = "Follow" // Sostituisci con il nome effettivo della tua tabella
	_, erro = db.Exec("DROP TABLE IF EXISTS " + tableName)
	if erro != nil {
		fmt.Println(erro)
		return nil, erro
	}

	//controllare se va qui
	_, err := db.Exec("PRAGMA foreign_keys = ON;")
	if err != nil {
		fmt.Println(err)
		return nil, ErrWithForeignKey
	}

	var User string
	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='User';`).Scan(&User)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE IF NOT EXISTS User (id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, username TEXT UNIQUE NOT NULL);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure User: %w", err)
		}
	}

	var Profile string
	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='Profile';`).Scan(&Profile)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE IF NOT EXISTS Profile (id INTEGER NOT NULL PRIMARY KEY, user TEXT NOT NULL, follower INTEGER, following INTEGER
			nPhoto INTEGER, FOREIGN KEY (id) REFERENCES User(id), FOREIGN KEY (user) REFERENCES User(username));`
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

	// Verifica se le chiavi esterne sono abilitate
	rows, errF := db.Query("PRAGMA foreign_keys;")
	if errF != nil {
		fmt.Println(errF)
		return nil, errF
	}
	var foreignKeysEnabled int
	for rows.Next() {
		err := rows.Scan(&foreignKeysEnabled)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	}
	if foreignKeysEnabled == 1 {
		log.Print("Le chiavi esterne sono abilitate.")
	} else {
		log.Print("Le chiavi esterne non sono abilitate.")
	}

	return &appdbimpl{
		c: db,
	}, nil
}

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}
