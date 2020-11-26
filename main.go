package main

import (
	"goblog/app/http/middlewares"
	"goblog/bootstrap"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var router *mux.Router

func main() {
	bootstrap.SetupDB()
	router = bootstrap.SetupRoute()

	http.ListenAndServe(":3030", middlewares.RemoveTrailingSlash(router))
}
