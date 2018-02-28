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
	CreateList(todo.List{Name: "List 1", Tasks: []todo.Task{{Action: "Buy tea powder"}}})
	CreateList(todo.List{Name: "List 2"})

	lists := GetLists()

	if len(lists) != 2 {
		t.Error()
	}
	if lists[0].Name != "List 1" || len(lists[0].Tasks) != 1 || lists[1].Name != "List 2" {
		t.Error()
	}
}

func Test_create_a_list(t *testing.T) {
	tasks := []todo.Task{{Action: "Buy tea powder"}, {Action: "Buy Mazaa"}}

	list := CreateList(todo.List{Name: "Shopping List", Tasks: tasks})

	if list.Name != "Shopping List" || list.Identifier == "" {
		t.Error()
	}

	for _, task := range list.Tasks {
		if task.Identifier == "" || task.ListIdentifier != list.Identifier {
			t.Error()
		}
	}
}
