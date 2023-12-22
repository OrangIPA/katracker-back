package db

import (
	"os"

	"github.com/jmoiron/sqlx"
)

var Conn *sqlx.DB

func ConnectDB() {
	Conn = sqlx.MustConnect("postgres", os.Getenv("db"))
}
