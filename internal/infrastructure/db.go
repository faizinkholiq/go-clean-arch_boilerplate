package infrastructure

import (
	"database/sql"
	"fmt"

	"github.com/faizinkholiq/gofiber_boilerplate/config"

	_ "github.com/lib/pq"
)

func InitDB() (*sql.DB, error) {
	config.LoadConfig()

	dbUser := config.GetEnv("DB_USER", "postgres")
	dbPass := config.GetEnv("DB_PASSWORD", "password")
	dbName := config.GetEnv("DB_NAME", "mydb")
	dbHost := config.GetEnv("DB_HOST", "localhost")

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=disable", dbUser, dbPass, dbName, dbHost)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}
