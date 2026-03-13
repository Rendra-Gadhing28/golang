package main

import (
	_"endpoint/crud"
	"endpoint/db"
	_"endpoint/handler"
	"endpoint/rest"
)

func main(){
	db.Database()
	rest.SetupRoute()
}