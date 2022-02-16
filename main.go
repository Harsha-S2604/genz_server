package main

import (

	"genz_server/routes"
	"genz_server/config/db"
)

func main() {
	/*
		1. database connection(MongoDB)
		2. setup router(Gin)
		3. start server
	*/
	database := db.ConnectDB()
	r := routes.SetupRouter(database)
	r.Run(":8081")
	
}