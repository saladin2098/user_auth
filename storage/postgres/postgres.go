package postgres

import (
	"database/sql"
	"fmt"

	"github.com/Mubinabd/auth_service/config"
	"github.com/Mubinabd/auth_service/storage"
	_ "github.com/lib/pq"
)

type Storage struct {
	db   *sql.DB
	UserS storage.UserI
}

func ConnectDB() (*Storage, error) {
	cfg := config.Load()
	dbConn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDatabase)
	db, err := sql.Open("postgres", dbConn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
    userS := NewUserStorage(db)
	return &Storage{
		db:      db,
		UserS: userS,
	}, nil
}
func (s *Storage) User() storage.UserI {
	if s.UserS == nil {
		s.UserS = NewUserStorage(s.db)
	}
	return s.UserS
}
