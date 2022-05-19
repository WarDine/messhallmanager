package submarineAPI

import (
	"repositories"
	"usecases"

	restful "github.com/emicklei/go-restful"
)

// NewSubmarineAPI returns an instance of a rest api handler
func NewMessHallManagerAPI() *restful.WebService {
	service := new(restful.WebService)
	service.Path("/messhallmanager").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	service.Route(service.GET("/messhall/info/").To(GetMessHallInfo))
	service.Route(service.POST("/messhall/add").To(AddMessHall))
	// service.Route(service.POST("/messhall/update").To(GetMessHallInfo))

	return service
}

func GetMessHallInfo(request *restful.Request, response *restful.Response) {
	messHall, err := repositories.PostgresRepo.GetAllMessHallsInfo()
	if err != nil {
		return
	}

	err = response.WriteEntity(messHall)
	if err != nil {
		return
	}
}

func AddMessHall(request *restful.Request, response *restful.Response) {
	err := repositories.PostgresRepo.AddMessHall(&usecases.MessHall{
		MessHallUID:      1234,
		Street:           "Rahova",
		City:             "Bucharest",
		Country:          "Romania",
		MenuUID:          321123,
		AttendanceNumber: 123,
	})

	// err := response.WriteEntity("Hello World!")
	if err != nil {
		return
	}
}
