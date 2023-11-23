package database

import (
	"database/sql"
	"errors"
	"fmt"

	api "github.com/Beatrice0399/WASAPhoto/service/api/reqcontext"
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
	Login(username string) error
	MyProfile() (api.Profile, error)
	SetMyUsername(id uint64, name string) error
	UploadPhoto(image []byte) (api.Photo, error)
	FollowUser() error
	UnfollowUser() error
	BanUser() error
	BannedUser() ([]api.Profile, error)
	UnbanUser() error
	GetUserProfile(id uint64) (api.Profile, error)
	GetMyStream(myId uint64) (api.Stream, error)
	LikePhoto(id uint64) error
	UnlikePhoto(id uint64) error
	CommentPhoto(id uint64, text string) (api.Comment, error)
	UncommentPhoto(id uint64) error
	DeletePhoto(id uint64) error

	Ping() error
}

var ErrProfileDoesNotExist = errors.New("Profile doesn't exist")

type appdbimpl struct {
	c *sql.DB
}

func New(db *sql.DB) (AppDatabase, error) {
	if db == nil {
		return nil, errors.New("Database is required whem building a AppDatabase")
	}
	var tableName string
	err := db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='Profile';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE Profile (id INTEGER NOT NULL PRIMARY KEY, name TEXT);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}

	return &appdbimpl{
		c: db,
	}, nil
}

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}
