package submarineAPI

import (
	"domain"
	"repositories"
	"usecases"

	restful "github.com/emicklei/go-restful"
	"github.com/google/uuid"
)

// NewSubmarineAPI returns an instance of a rest api handler
func NewMessHallManagerAPI() *restful.WebService {
	cors := restful.CrossOriginResourceSharing{
		// ExposeHeaders:  []string{"X-My-Header"},
		AllowedMethods: []string{"GET", "POST", "DELETE", "UPDATE"},
		AllowedHeaders: []string{"Content-Type", "Accept", "Authorization"},
		AllowedDomains: []string{"*"},
		CookiesAllowed: false,
	}

	service := new(restful.WebService)
	service.Filter(cors.Filter)

	service.Path("/messhallmanager").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	service.Route(service.GET("/messhall/info/").To(GetMessHallInfo))
	service.Route(service.GET("/messhall/info/{messhall_id}").To(GetMessHallInfoByID))
	service.Route(service.GET("/messhall/menu/{messhall_id}").To(GetMessHallMenu))

	service.Route(service.POST("/messhall/checkin/{messhall_id}").To(CheckinHandler))
	service.Route(service.POST("/messhall/checkout/{messhall_id}").To(CheckoutHandler))

	service.Route(service.GET("/messhalladmin/info/").To(GetMessHallAdminInfo))
	service.Route(service.GET("/messhalladmin/info/{messhalladmin_id}").To(GetMessHallAdminInfoByID))

	service.Route(service.POST("/messhall/add").To(AddMessHall))
	service.Route(service.DELETE("/messhall/delete/{messhall_id}").To(DeleteMessHall))

	service.Route(service.POST("/messhall/update/{messhall_id}/{new_status}").To(UpdateMessHall))

	return service
}

// CheckinHandler
func CheckinHandler(request *restful.Request, response *restful.Response) {
	messhallID := request.PathParameter("messhall_id")

	err := repositories.PostgresRepo.IncrementMessHallAttendance(messhallID)
	if err != nil {
		err = response.WriteEntity("Checkin failed")
		if err != nil {
			return
		}
		return
	}

	err = response.WriteEntity("Checkin complete")
	if err != nil {
		return
	}
}

// CheckoutHandler
func CheckoutHandler(request *restful.Request, response *restful.Response) {
	messhallID := request.PathParameter("messhall_id")

	err := repositories.PostgresRepo.DecrementMessHallAttendance(messhallID)
	if err != nil {
		err = response.WriteEntity("Checkout failed")
		if err != nil {
			return
		}
		return
	}

	err = response.WriteEntity("Checkout complete")
	if err != nil {
		return
	}
}

// GetMessHallMenu
func GetMessHallMenu(request *restful.Request, response *restful.Response) {
	messhallID := request.PathParameter("messhall_id")

	messHallMenuEntries, err := repositories.PostgresRepo.GetMessHallMenuInfo(messhallID)
	if err != nil {
		return
	}

	menuRecipes := []usecases.Recipe{}
	for _, menuEntry := range messHallMenuEntries {

		recipe := repositories.PostgresRepo.GetRecipesByRecipeUID(menuEntry.RecipeUID)
		menuRecipes = append(menuRecipes, recipe[0])
	}

	err = response.WriteEntity(menuRecipes)
	if err != nil {
		return
	}
}

// GetMessHallInfo returns info about all mess halls
func GetMessHallInfo(request *restful.Request, response *restful.Response) {
	messHalls, err := repositories.PostgresRepo.GetAllMessHallsInfo()
	if err != nil {
		return
	}

	err = response.WriteEntity(messHalls)
	if err != nil {
		return
	}
}

// GetMessHallInfoByID
func GetMessHallInfoByID(request *restful.Request, response *restful.Response) {
	messhallID := request.PathParameter("messhall_id")

	messHall, err := repositories.PostgresRepo.GetMessHallsInfoByID(messhallID)
	if err != nil {
		return
	}

	err = response.WriteEntity(messHall)
	if err != nil {
		return
	}
}

// GetMessHallAdminInfo
func GetMessHallAdminInfo(request *restful.Request, response *restful.Response) {
	messHallAdmins, err := repositories.PostgresRepo.GetAllMessHallAdminsInfo()
	if err != nil {
		return
	}

	err = response.WriteEntity(messHallAdmins)
	if err != nil {
		return
	}
}

// GetMessHallAdminInfoByID
func GetMessHallAdminInfoByID(request *restful.Request, response *restful.Response) {
	messhallAdminID := request.PathParameter("messhalladmin_id")

	messHallAdmin, err := repositories.PostgresRepo.GetMessHallAdminsInfoByID(messhallAdminID)
	if err != nil {
		return
	}

	err = response.WriteEntity(messHallAdmin)
	if err != nil {
		return
	}
}

// AddMessHall add a new mess hall and its admin
func AddMessHall(request *restful.Request, response *restful.Response) {
	var queryBody domain.AddMessHallInfoQuery

	/* Get mess info from query */
	err := request.ReadEntity(&queryBody)
	if err != nil {
		err = response.WriteEntity("ERROR: new mess hall data mallformed")
		if err != nil {
			return
		}
		return
	}

	/* Populate new mess hall struct */
	messHallUID := uuid.New().String()
	newMessHall := usecases.MessHall{
		MessHallUID:      messHallUID,
		Street:           queryBody.Street,
		City:             queryBody.City,
		Country:          queryBody.Country,
		Status:           queryBody.Status,
		AttendanceNumber: queryBody.AttendanceNumber,
	}

	/* Populat new mess hall admin struct */
	newMessHallAdmin := usecases.MessHallAdmin{
		MessHallAdminUID: uuid.New().String(),
		Nickname:         queryBody.MessHallAdminNickname,
		MessHallUID:      messHallUID,
	}

	/* Add mess hall info and its admin info to repository */
	err = repositories.PostgresRepo.AddMessHall(&newMessHall, &newMessHallAdmin)
	if err != nil {
		err = response.WriteEntity("ERROR: failed to add new mess hall")
		if err != nil {
			return
		}
		return
	}

	err = response.WriteEntity("SUCCESS: added new mess hall info")
	if err != nil {
		return
	}
}

// UpdateMessHall
func UpdateMessHall(request *restful.Request, response *restful.Response) {
	messhallID := request.PathParameter("messhall_id")
	newStatus := request.PathParameter("new_status")

	err := repositories.PostgresRepo.UpdateMessHallStatus(messhallID, newStatus)
	if err != nil {
		err = response.WriteEntity("ERROR: failed to update mess hall status")
		if err != nil {
			return
		}
		return
	}

	err = response.WriteEntity("SUCCESS: mess hall status update.")
	if err != nil {
		return
	}
}

// DeleteMessHall
func DeleteMessHall(request *restful.Request, response *restful.Response) {
	messhallID := request.PathParameter("messhall_id")

	err := repositories.PostgresRepo.DeleteMessHall(messhallID)
	if err != nil {
		err = response.WriteEntity("ERROR: failed to delete mess hall")
		if err != nil {
			return
		}
		return
	}

	err = response.WriteEntity("SUCCESS: deleted mess hall.")
	if err != nil {
		return
	}
}
