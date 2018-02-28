package server

import (
	"github.com/aj-jaswanth/sample/src/server/handler"
	"github.com/aj-jaswanth/sample/src/util"
	"net/http"
)

func Serve() {
	http.HandleFunc("/lists", handler.List)
	err := http.ListenAndServe("localhost:8080", nil)
	util.HandleErr(err)
}
