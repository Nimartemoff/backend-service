package main

import (
	"BackendService/app/database"
	"BackendService/app/routes"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	dbconn, err := database.CreateDBConnection()
	if err != nil {
		fmt.Println("Failed to connect to the database: ", err)
		return
	}

	defer dbconn.Close()
	router := routes.SetupRouter(dbconn)
	router.Run(":8080")
}
