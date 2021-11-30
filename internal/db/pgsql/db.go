package pgsql

import (
	"fmt"
	"log"

	// Sqlx
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/spf13/viper"
)

var (
	// Sqlx
	Client *sqlx.DB

	// conf
	username string
	password string
	host string
	db string
	Schema string
	port string
)

func init() {
	var err error

	// viper
	viper.AddConfigPath("/app/internal/conf")
    viper.SetConfigName("config")
    viper.SetConfigType("json")

	errr := viper.ReadInConfig()
    if err != nil {
		panic(errr)
    }
	
	// passport
	username = viper.GetString("pgsql.username")
	password = viper.GetString("pgsql.password")
	host = viper.GetString("pgsql.host")
	db = viper.GetString("pgsql.db")
	Schema = viper.GetString("pgsql.schema")
	port = viper.GetString("pgsql.port")

	// conection string
	dataSourceName := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s search_path=%s sslmode=disable", host, username, password, db, port, Schema)
	
	Client, err = sqlx.Connect("postgres", dataSourceName)
	if err != nil {
		panic(err)
	}

	log.Println("database successfully configured")
}
