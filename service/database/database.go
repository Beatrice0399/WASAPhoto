package database

import (
	"database/sql"
	"fmt"

	api "github.com/Beatrice0399/WASAPhoto/service/api/reqcontext"
)

type DB struct {
	Filename string `conf:""`
}

// DA TOGLIERE? Open IS USED TO CREATE A DB HANDLE
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

type AppDatabase interface {
	Login(username string) error
	MyProfile() (api.Profile, error)
	SetMyUsername(api.User) error
	UploadPhoto(image []byte) (api.Photo, error)
	FollowUser() error
	UnfollowUser() error
	BanUser() error
	UnbanUser() error
	GetUserProfile(id uint64) (api.Profile, error)
	GetMyStream(myId uint64) (api.Stream, error)
	LikePhoto(id uint64) error
	UnlikePhoto(id uint64) error
	CommentPhoto(id uint64, text string) (api.Comment, error)
	UncommentPhoto(id uint64) error
	DeletePhoto(id uint64) error
}

type appdbimpl struct {
	c *sql.DB
}
