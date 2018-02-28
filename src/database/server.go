package database

import (
	"database/sql"
	"github.com/aj-jaswanth/sample/src/util"
	_ "github.com/lib/pq"
)

func AcquireDatabase() *sql.DB {
	db, err := sql.Open("postgres", "postgres://postgres@localhost/todo?sslmode=disable")
	util.HandleErr(err)
	err = db.Ping()
	util.HandleErr(err)
	return db
}
