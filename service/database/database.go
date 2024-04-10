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

// Open IS USED TO CREATE A DB HANDLE

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
var ErrProfileDoesNotExist = errors.New("Profile doesn't exist")
var ErrUsernameUsed = errors.New("Username already used")
var ErrFollowUser = errors.New("This user banned you")
var ErrWithForeignKey = errors.New("Error turning on foreign key")
var ErrUserBanned = errors.New("You banned this user")
var ErrUserBannedYou = errors.New("This user banned you")
var ErrPhoto = errors.New("This photo doesn't exsist")
var ErrLike = errors.New("Inconsistent like id")

type AppDatabase interface {
	DoLogin(username string) (int, error)
	SearchUser(myId int, username string) ([]User, error)
	SetMyUsername(id int, name string) error
	NewPhoto(id int) (int, error)
	UpdatePathPhoto(phid int, path string) error
	GetPhoto(phId int) (Photo, error)
	FollowUser(myId int, fid int) error
	UnfollowUser(myId int, fid int) error
	BanUser(myId int, bid int) error
	BannedUser(myId int) ([]User, error)
	UnbanUser(myId int, bid int) error
	GetUserProfile(id int, myId int) (Profile, error)
	GetMyStream(myId int) ([]Photo, error)
	LikePhoto(phId int, uid int) error
	UnlikePhoto(id int, uid int, lid int) error
	CommentPhoto(uid int, phid int, text string) (int, error)
	UncommentPhoto(cid int, phid int, uid int) error
	DeletePhoto(phid int, myid int) error

	IsBanned(myId int, idProfile int) bool
	GetPhotoUser(id int) ([]Photo, error)
	GetPhotoComments(phId int) ([]Comment, error)
	GetFollower(id int) ([]User, error)
	GetFollowing(followedBy int) ([]User, error)
	GetId(username string) (int, error)
	GetNameById(id int) (string, error)
	GetLikesPhoto(phid int) ([]User, error)

	Ping() error
}

type appdbimpl struct {
	c *sql.DB
}

func New(db *sql.DB) (AppDatabase, error) {
	if db == nil {
		return nil, errors.New("Database is required whem building a AppDatabase")
	}

	// DROP TABLE
	/*
		tableName := "Follow" // Sostituisci con il nome effettivo della tua tabella
		_, erro := db.Exec("DROP TABLE IF EXISTS " + tableName)
		if erro != nil {
			fmt.Println(erro)
			return nil, erro
		}
		tableName = "Ban" // Sostituisci con il nome effettivo della tua tabella
		_, erro = db.Exec("DROP TABLE IF EXISTS " + tableName)
		if erro != nil {
			fmt.Println(erro)
			return nil, erro
		}
		tableName = "Likes" // Sostituisci con il nome effettivo della tua tabella
		_, erro = db.Exec("DROP TABLE IF EXISTS " + tableName)
		if erro != nil {
			fmt.Println(erro)
			return nil, erro
		}
		tableName = "Comment" // Sostituisci con il nome effettivo della tua tabella
		_, erro = db.Exec("DROP TABLE IF EXISTS " + tableName)
		if erro != nil {
			fmt.Println(erro)
			return nil, erro
		}
		tableName = "Photo" // Sostituisci con il nome effettivo della tua tabella
		_, erro = db.Exec("DROP TABLE IF EXISTS " + tableName)
		if erro != nil {
			fmt.Println(erro)
			return nil, erro
		}
		tableName = "User" // Sostituisci con il nome effettivo della tua tabella
		_, erro = db.Exec("DROP TABLE IF EXISTS " + tableName)
		if erro != nil {
			fmt.Println(erro)
			return nil, erro
		}
	*/
	_, err := db.Exec("PRAGMA foreign_keys = ON;")
	if err != nil {
		return nil, ErrWithForeignKey
	}

	var User string
	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='User';`).Scan(&User)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE IF NOT EXISTS User (id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, username VARCHAR(16) UNIQUE NOT NULL);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure User: %w", err)
		}
	}

	var Photo string
	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='Photo';`).Scan(&Photo)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE IF NOT EXISTS Photo (id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, user INTEGER NOT NULL, image_path TEXT,
			date DATETIME NOT NULL,
			FOREIGN KEY (user) REFERENCES User(id) ON DELETE CASCADE);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure Photo: %w", err)
		}
	}

	var Comment string
	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='Comment';`).Scan(&Comment)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE IF NOT EXISTS Comment (id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, user INTEGER NOT NULL, photo INTEGER NOT NULL, string TEXT NOT NULL, 
			date DATETIME NOT NULL, FOREIGN KEY (user) REFERENCES User(id) ON DELETE CASCADE, FOREIGN KEY (photo) REFERENCES Photo(id) ON DELETE CASCADE);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure Comment: %w", err)
		}
	}

	var Follow string
	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='Follow';`).Scan(&Follow)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE IF NOT EXISTS Follow (user INTEGER NOT NULL, followedBy INTEGER NOT NULL,
			UNIQUE(user, followedBy), FOREIGN KEY (user) REFERENCES User(id), FOREIGN KEY (followedBy) REFERENCES User(id));`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure Follow: %w", err)
		}
	}

	var Ban string
	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='Ban';`).Scan(&Ban)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE IF NOT EXISTS Ban (banned INTEGER NOT NULL, whoBan INTEGER NOT NULL,
			UNIQUE(banned, whoBan), FOREIGN KEY (banned) REFERENCES User(id), FOREIGN KEY (whoBan) REFERENCES User(id));`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure Ban: %w", err)
		}
	}

	var Likes string
	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='Likes';`).Scan(&Likes)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE IF NOT EXISTS Likes (phId INTEGER NOT NULL, uid INTEGER NOT NULL,
					UNIQUE(phId, uid), FOREIGN KEY (phId) REFERENCES Photo(id) ON DELETE CASCADE, FOREIGN KEY (uid) REFERENCES User(id) ON DELETE CASCADE);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure Likes: %w", err)
		}
	}

	// Verifica se le chiavi esterne sono abilitate
	/*
		rows, errF := db.Query("PRAGMA foreign_keys;")
		if errF != nil {
			// fmt.Println(errF)
			return nil, errF
		}
		var foreignKeysEnabled int
		for rows.Next() {
			err := rows.Scan(&foreignKeysEnabled)
			if err != nil {
				// fmt.Println(err)
				return nil, err
			}
		}
		if foreignKeysEnabled == 1 {
			log.Print("Le chiavi esterne sono abilitate.")
		} else {
			log.Print("Le chiavi esterne non sono abilitate.")
		}
	*/

	return &appdbimpl{
		c: db,
	}, nil
}

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}
