package handler

import (
	"encoding/json"
	"fmt"
	"github.com/aj-jaswanth/sample/src/todo"
	"github.com/aj-jaswanth/sample/src/todo/persistence"
	"github.com/aj-jaswanth/sample/src/util"
	"io/ioutil"
	"net/http"
)

func List(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		createAList(w, r)
	} else {
		getAllLists(w)
	}
}

func createAList(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	list := new(todo.List)
	jsonBytes, err := ioutil.ReadAll(r.Body)
	util.HandleErr(err)
	err = json.Unmarshal(jsonBytes, list)
	util.HandleErr(err)
	persistence.CreateList(*list)
}

func getAllLists(w http.ResponseWriter) {
	lists := persistence.GetLists()
	jsonBytes, err := json.Marshal(lists)
	util.HandleErr(err)
	fmt.Fprintln(w, string(jsonBytes))
}
