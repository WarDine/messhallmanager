package main

import (
	// "events"
	//"events"
	"log"
	"net/http"
	"repositories"
	messHallManagerAPI "submarineapi"

	"github.com/emicklei/go-restful"
)

func main() {
	// Init Repositories
	repositories.PostgresRepo = *repositories.NewPostgresManager()

	// Start http server
	restful.Add(messHallManagerAPI.NewMessHallManagerAPI())
	log.Fatal(http.ListenAndServe(":8090", nil))

}
