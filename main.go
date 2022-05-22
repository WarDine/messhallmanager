package main

import (
	// "events"
	//"events"

	"repositories"
	messHallManagerAPI "submarineapi"
)

func main() {
	// Init Repositories
	repositories.PostgresRepo = *repositories.NewPostgresManager()

	// Start http server
	// restful.Add(messHallManagerAPI.NewMessHallManagerAPI())
	// log.Fatal(http.ListenAndServe(":8090", nil))
	messHallManagerAPI.StartServer()

}
