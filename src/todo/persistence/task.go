package persistence

import (
	"github.com/aj-jaswanth/sample/src/database"
	"github.com/aj-jaswanth/sample/src/todo"
	"github.com/aj-jaswanth/sample/src/util"
)

func GetTasks(listIdentifier string) []todo.Task {
	db := database.AcquireDatabase()
	rows, err := db.Query("select * from task where list_identifier=$1", listIdentifier)
	util.HandleErr(err)
	tasks := make([]todo.Task, 0)
	for rows.Next() {
		task := todo.Task{}
		err := rows.Scan(&task.Identifier, &task.Action, &task.ListIdentifier)
		util.HandleErr(err)
		tasks = append(tasks, task)
	}
	return tasks
}

func CreateTask(task *todo.Task) *todo.Task {
	task.Identifier = util.NewUUID()
	db := database.AcquireDatabase()
	_, err := db.Exec("INSERT INTO task VALUES ($1, $2, $3)", task.Identifier, task.Action, task.ListIdentifier)
	util.HandleErr(err)
	return task
}
