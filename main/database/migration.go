package database

import (
	"fmt"
	"github.com/aj-jaswanth/sample/main/util"
	"log"
)

var migrations = []string{
	`create table _migrations
		(index integer)`,
	`create table list 
		(identifier uuid primary key, 
		name varchar unique not null)`,
	`create table task 
		(identifier uuid primary key,
		content varchar,
		list_identifier uuid references list(identifier))`,
}

func shouldRun(index int) bool {
	db := AcquireDatabase()
	row := db.QueryRow("select * from _migrations where index=$1", index)
	var dummy int
	err := row.Scan(&dummy)
	if err != nil {
		log.Printf("Error accessing migrations table: %s", err)
		return true
	}
	return false
}

func RunMigrations() {
	db := AcquireDatabase()
	for index, migration := range migrations {
		if shouldRun(index) {
			fmt.Printf("Running migration %s\n", migration)
			_, err := db.Exec(migration)
			if err == nil {
				db.Exec("insert into _migrations values($1)", index)
			}
			util.HandleErr(err)
		}
	}
}

func Erase() {
	db := AcquireDatabase()
	_, err := db.Exec("drop table _migrations")
	util.HandleErr(err)
	_, err = db.Exec("drop table task")
	util.HandleErr(err)
	_, err = db.Exec("drop table list")
	util.HandleErr(err)
}
