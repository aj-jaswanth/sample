package persistence

import (
	"github.com/aj-jaswanth/sample/src/database"
	"github.com/aj-jaswanth/sample/src/todo"
	"github.com/aj-jaswanth/sample/src/util"
)

func GetLists() []todo.List {
	db := database.AcquireDatabase()
	rows, err := db.Query("select * from list")
	util.HandleErr(err)
	var lists []todo.List
	for rows.Next() {
		list := todo.List{}
		err := rows.Scan(&list.Identifier, &list.Name)
		util.HandleErr(err)
		list.Tasks = GetTasks(list.Identifier)
		lists = append(lists, list)
	}
	return lists
}

func CreateList(list todo.List) todo.List {
	db := database.AcquireDatabase()
	list.Identifier = util.NewUUID()
	_, err := db.Exec("INSERT INTO list VALUES ($1, $2)", list.Identifier, list.Name)
	if list.Tasks != nil {
		for idx := range list.Tasks {
			task := &list.Tasks[idx]
			task.ListIdentifier = list.Identifier
			CreateTask(task)
		}
	}
	util.HandleErr(err)
	return list
}
