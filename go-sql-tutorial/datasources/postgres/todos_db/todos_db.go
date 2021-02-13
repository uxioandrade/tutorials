package todos_db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/stdlib"
)

const (
	dbUsername = "USERNAME"
	dbPassword = "PASSWORD"
	dbHost     = "HOST"
	dbSchema   = "SCHEMA"
)

var (
	Client   *sql.DB
	username = os.Getenv(dbUsername)
	password = os.Getenv(dbPassword)
	host     = os.Getenv(dbHost)
	schema   = os.Getenv(dbSchema)
)

func init() {
	connInfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		host, username, password, schema)
	var err error
	Client, err = sql.Open("pgx", connInfo)
	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("Database ready to accept connections")
}
