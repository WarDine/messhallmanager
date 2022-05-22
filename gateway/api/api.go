package submarineAPI

import (
	"domain"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"repositories"
	"usecases"

	"github.com/google/uuid"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

const (
	HttpServerPort = ":8081"
)

func NewRecipeAPI() *mux.Router {

	var router = mux.NewRouter()
	router.Use(commonMiddleware)
	router.Use(mux.CORSMethodMiddleware(router))

	router.HandleFunc("/messhallmanager/messhall/info", GetMessHallInfo).Methods("GET")
	router.HandleFunc("/messhallmanager/messhall/info/{messhall_id}", GetMessHallInfoByID).Methods("GET")
	router.HandleFunc("/messhallmanager/messhall/menu/{messhall_id}", GetMessHallMenu).Methods("GET")

	router.HandleFunc("/messhallmanager/messhall/checkin/{messhall_id}", CheckinHandler).Methods("POST")
	router.HandleFunc("/messhallmanager/messhall/checkout/{messhall_id}", CheckoutHandler).Methods("POST")

	router.HandleFunc("/messhallmanager/messhalladmin/info", GetMessHallAdminInfo).Methods("GET")
	router.HandleFunc("/messhallmanager/messhalladmin/info/{messhalladmin_id}", GetMessHallAdminInfoByID).Methods("GET")

	router.HandleFunc("/messhallmanager/messhall/add", AddMessHall).Methods("POST")
	router.HandleFunc("/messhallmanager/messhall/add", optionsHandler).Methods("OPTIONS")

	router.HandleFunc("/messhallmanager/messhall/delete/{messhall_id}", DeleteMessHall).Methods("DELETE")

	router.HandleFunc("/messhallmanager/messhall/update/{messhall_id}/{new_status}", UpdateMessHall).Methods("POST")
	router.HandleFunc("/messhallmanager/messhall/update", optionsHandler).Methods("OPTIONS")

	return router
}

func StartServer() {
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization", "Origin", "application/json"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "UPDATE", "OPTIONS", "PUT", "PATCH"})

	router := NewRecipeAPI()

	fmt.Printf("HTTP Server is running at http://localhost%s\n", HttpServerPort)
	log.Fatal(http.ListenAndServe(HttpServerPort, handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

func optionsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("!!!!!!!!!!!!!!!!!!!!REceived options request!!!!!!!!!!!!!!!!!!!!")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode("Received options. Send 200 ok")
	if err != nil {
		return
	}
	return
}

// CheckinHandler
func CheckinHandler(w http.ResponseWriter, r *http.Request) { //(request *restful.Request, response *restful.Response) {
	params := mux.Vars(r)
	messhallID := params["messhall_id"]

	err := repositories.PostgresRepo.IncrementMessHallAttendance(messhallID)
	if err != nil {
		err = json.NewEncoder(w).Encode("Checkin failed")
		if err != nil {
			return
		}
		return
	}

	err = json.NewEncoder(w).Encode("Checkin complete")
	if err != nil {
		return
	}
}

// CheckoutHandler
func CheckoutHandler(w http.ResponseWriter, r *http.Request) { //(request *restful.Request, response *restful.Response) {
	params := mux.Vars(r)
	messhallID := params["messhall_id"]

	err := repositories.PostgresRepo.DecrementMessHallAttendance(messhallID)
	if err != nil {
		err = json.NewEncoder(w).Encode("Checkout failed")
		if err != nil {
			return
		}
		return
	}

	err = json.NewEncoder(w).Encode("Checkout complete")
	if err != nil {
		return
	}
}

// GetMessHallMenu
func GetMessHallMenu(w http.ResponseWriter, r *http.Request) { //(request *restful.Request, response *restful.Response) {
	params := mux.Vars(r)
	messhallID := params["messhall_id"]

	messHallMenuEntries, err := repositories.PostgresRepo.GetMessHallMenuInfo(messhallID)
	if err != nil {
		return
	}

	menuRecipes := []usecases.Recipe{}
	for _, menuEntry := range messHallMenuEntries {

		recipe := repositories.PostgresRepo.GetRecipesByRecipeUID(menuEntry.RecipeUID)
		menuRecipes = append(menuRecipes, recipe[0])
	}

	err = json.NewEncoder(w).Encode(menuRecipes)
	if err != nil {
		return
	}
}

// GetMessHallInfo returns info about all mess halls
func GetMessHallInfo(w http.ResponseWriter, r *http.Request) { //(request *restful.Request, response *restful.Response) {
	messHalls, err := repositories.PostgresRepo.GetAllMessHallsInfo()
	if err != nil {
		return
	}

	err = json.NewEncoder(w).Encode(messHalls)
	if err != nil {
		return
	}
}

// GetMessHallInfoByID
func GetMessHallInfoByID(w http.ResponseWriter, r *http.Request) { //(request *restful.Request, response *restful.Response) {
	params := mux.Vars(r)
	messhallID := params["messhall_id"]

	messHall, err := repositories.PostgresRepo.GetMessHallsInfoByID(messhallID)
	if err != nil {
		return
	}

	err = json.NewEncoder(w).Encode(messHall)
	if err != nil {
		return
	}
}

// GetMessHallAdminInfo
func GetMessHallAdminInfo(w http.ResponseWriter, r *http.Request) { //(request *restful.Request, response *restful.Response) {
	messHallAdmins, err := repositories.PostgresRepo.GetAllMessHallAdminsInfo()
	if err != nil {
		return
	}

	err = json.NewEncoder(w).Encode(messHallAdmins)
	if err != nil {
		return
	}
}

// GetMessHallAdminInfoByID
func GetMessHallAdminInfoByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	messhallAdminID := params["messhalladmin_id"]

	messHallAdmin, err := repositories.PostgresRepo.GetMessHallAdminsInfoByID(messhallAdminID)
	if err != nil {
		return
	}

	err = json.NewEncoder(w).Encode(messHallAdmin)
	if err != nil {
		return
	}
}

// AddMessHall add a new mess hall and its admin
func AddMessHall(w http.ResponseWriter, r *http.Request) {
	var queryBody domain.AddMessHallInfoQuery

	/* Get mess info from query */
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&queryBody)
	if err != nil {
		err = json.NewEncoder(w).Encode("ERROR: new mess hall data mallformed")
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
		err = json.NewEncoder(w).Encode("ERROR: failed to add new mess hall")
		if err != nil {
			return
		}
		return
	}
	err = json.NewEncoder(w).Encode("SUCCESS: added new mess hall info")
	if err != nil {
		return
	}
}

func UpdateMessHall(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	messhallID := params["messhall_id"]
	newStatus := params["new_status"]

	err := repositories.PostgresRepo.UpdateMessHallStatus(messhallID, newStatus)
	if err != nil {
		err = json.NewEncoder(w).Encode("ERROR: failed to update mess hall status")
		if err != nil {
			return
		}
		return
	}
	err = json.NewEncoder(w).Encode("SUCCESS: mess hall status update.")
	if err != nil {
		return
	}
}

// DeleteMessHall
func DeleteMessHall(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	messhallID := params["messhall_id"]

	err := repositories.PostgresRepo.DeleteMessHall(messhallID)
	if err != nil {
		err = json.NewEncoder(w).Encode("ERROR: failed to delete mess hall")
		if err != nil {
			return
		}
		return
	}

	err = json.NewEncoder(w).Encode("SUCCESS: deleted mess hall.")
	if err != nil {
		return
	}
}
