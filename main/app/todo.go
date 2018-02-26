package app

import (
	"github.com/aj-jaswanth/sample/main/database"
	"github.com/aj-jaswanth/sample/main/util"
)

type Task struct {
	Identifier     string
	Content        string
	ListIdentifier string
}

type List struct {
	Identifier string
	Name       string
}

func CreateList(list List) List {
	db := database.AcquireDatabase()
	list.Identifier = util.NewUUID()
	_, err := db.Exec("INSERT INTO list VALUES ($1, $2)", list.Identifier, list.Name)
	util.HandleErr(err)
	return list
}

func CreateTask(task Task) Task {
	task.Identifier = util.NewUUID()
	db := database.AcquireDatabase()
	_, err := db.Exec("INSERT INTO task VALUES ($1, $2, $3)", task.Identifier, task.Content, task.ListIdentifier)
	util.HandleErr(err)
	return task
}

func AddTaskTo(task Task, listName string) Task {
	db := database.AcquireDatabase()
	row := db.QueryRow("SELECT * FROM list WHERE name=$1", listName)
	list := new(List)
	err := row.Scan(&list.Identifier, &list.Name)
	util.HandleErr(err)
	task.ListIdentifier = list.Identifier
	return CreateTask(task)
}
