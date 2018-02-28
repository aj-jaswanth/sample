package persistence

import (
	"github.com/aj-jaswanth/sample/src/database"
	"github.com/aj-jaswanth/sample/src/todo"
	"testing"
)

func Test_Run_Migrations(t *testing.T) {
	database.RunMigrations()
}

func Test_get_all_lists(t *testing.T) {
	database.Truncate()
	CreateList(todo.List{Name: "List 1"})
	CreateList(todo.List{Name: "List 2"})
}

func Test_create_a_list(t *testing.T) {
	tasks := []todo.Task{
		{Action: "Buy tea powder"}, {Action: "Buy Mazaa"}}

	list := CreateList(todo.List{Name: "Shopping List", Tasks: tasks})

	if list.Name != "Shopping List" || list.Identifier == "" {
		t.Fatal(list)
	}

	for _, task := range list.Tasks {
		if task.Identifier == "" || task.ListIdentifier != list.Identifier {
			t.Fatal(list)
		}
	}
}
