package app

import (
	"testing"
	"github.com/aj-jaswanth/sample/main/database"
)

func Test_Run_Migrations(t *testing.T) {
	database.RunMigrations()
}

func Test_Create_A_List(t *testing.T) {
	list := CreateList(List{Name: "Shopping List"})

	if list.Name != "Shopping List" || !(len(list.Identifier) > 0) {
		t.Fatal(list)
	}
}

func Test_Add_Task_to_List(t *testing.T) {
	task := AddTaskTo(Task{Content: "Buy food"}, "Shopping List")

	if !(len(task.Identifier) > 0 ) || task.Content != "Buy food" {
		t.Fatal(task)
	}
}

func Test_Erase_Db(t *testing.T) {
	database.Erase()
}
