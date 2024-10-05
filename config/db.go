package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func InitDB(cfg Config) (*sql.DB, error) {
	dbUser := cfg.DB.User
	dbPass := cfg.DB.Password
	dbName := cfg.DB.Name
	dbHost := cfg.DB.Host

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=disable", dbUser, dbPass, dbName, dbHost)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	return db, nil
}
