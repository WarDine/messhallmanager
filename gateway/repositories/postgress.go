package repositories

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
	"usecases"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	selectAllMessHallInfoQuery  = "SELECT * FROM messhall;"
	selectMessHallInfoByIDQuery = "SELECT * FROM messhall WHERE messhalls_uid = '%s';"
	selectMessHallMenuInfoQuery = "SELECT * FROM menu WHERE menu_uid = '%s';"

	selectAllMessHallAdminInfoQuery  = "SELECT * FROM messhalls_admins;"
	selectMessHallAdminInfoByIDQuery = "SELECT * FROM messhalls_admins WHERE messhalls_admins_uid = '%s';"

	updateMessHallStatusQuery = "UPDATE messhall SET status = $2 WHERE messhalls_uid = $1;"

	incrementMessHallAttendanceQuery = "UPDATE messhall SET attendance_number = attendance_number + 1 WHERE messhalls_uid = $1;"
	decrementMessHallAttendanceQuery = "UPDATE messhall SET attendance_number = attendance_number - 1 WHERE messhalls_uid = $1;"

	deleteMessHallQuery      = "DELETE FROM messhall WHERE messhalls_uid = $1;"
	deleteMessHallAdminQuery = "DELETE FROM messhalls_admins WHERE messhalls_admins.messhall_uid = $1;"

	getRecipesByIDQuery = "SELECT * FROM %s WHERE recipe_uid='%s';"
)

type PostgresManager struct {
	conn *sqlx.DB
}

var PostgresRepo PostgresManager

// NewPostgresManager return postgres objects which contains connection to database
func NewPostgresManager() *PostgresManager {

	port, err := strconv.Atoi(os.Getenv("PGPORT"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("PGPORT is: ", port)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("PGHOST"), port, os.Getenv("PGUSER"), os.Getenv("PGPASSWORD"), os.Getenv("PGDATABASE"))

	log.Println(".env: ", psqlInfo)

	time.Sleep(20 * time.Second)
	// conn, err := sql.Open("postgres", psqlInfo)
	conn, err := sqlx.Connect("postgres", psqlInfo)
	if err != nil {
		time.Sleep(20 * time.Second)
		conn, err = sqlx.Connect("postgres", psqlInfo)
		if err != nil {
			panic(err)
		}
	}

	return &PostgresManager{
		conn: conn,
	}
}

//DeleteConnection closes DB connection
func (pg *PostgresManager) DeleteConnection() {
	pg.conn.Close()
}

// getListDBMEssHallTags
func getListDBMEssHallTags(messHall *usecases.MessHall) []string {

	t := reflect.TypeOf(*messHall)

	tagFields := make([]string, t.NumField())
	for i := range tagFields {
		tagFields[i] = getDBTagName(messHall, t.Field(i).Name)
	}

	return tagFields
}

// getListDBMEssHallAdminTags
func getListDBMEssHallAdminTags(messHall *usecases.MessHallAdmin) []string {

	t := reflect.TypeOf(*messHall)

	tagFields := make([]string, t.NumField())
	for i := range tagFields {
		tagFields[i] = getDBTagName(messHall, t.Field(i).Name)
	}

	return tagFields
}

// getDBTagName returns a field from a generic struct
func getDBTagName(genericStruct interface{}, structField string) string {

	tagName := "db"
	field, ok := reflect.TypeOf(genericStruct).Elem().FieldByName(structField)
	if !ok {
		log.Fatal("Field not found")
	}

	return string(field.Tag.Get(tagName))
}

/**
 * createQueryFields generates a string like:
 * "ingredient_name, recipe_uid, amount"
 */
func createQueryFields(fields []string) string {

	var queryFields strings.Builder

	for i, field := range fields {
		if i == len(fields)-1 {
			queryFields.WriteString(field)
			break
		}

		queryFields.WriteString(field)
		queryFields.WriteString(", ")
	}

	return queryFields.String()
}

/**
 * createQueryValues generates a string like:
 * ":ingredient_name, :recipe_uid, :amount"
 */
func createQueryValues(fields []string) string {

	var queryFields strings.Builder

	for i, field := range fields {
		queryFields.WriteString(":")
		if i == len(fields)-1 {
			queryFields.WriteString(field)
			break
		}

		queryFields.WriteString(field)
		queryFields.WriteString(", ")
	}

	return queryFields.String()
}

// GetAllMessHallsInfo
func (pg *PostgresManager) GetMessHallsInfoByID(id string) ([]usecases.MessHall, error) {

	db := pg.conn

	messHalls := []usecases.MessHall{}
	query := fmt.Sprintf(selectMessHallInfoByIDQuery, id)
	err := db.Select(&messHalls, query)

	if err != nil {
		return nil, err
	}

	return messHalls, nil
}

// GetAllMessHallsInfo returns info about all Mess Halls
func (pg *PostgresManager) GetAllMessHallsInfo() ([]usecases.MessHall, error) {

	db := pg.conn

	messHalls := []usecases.MessHall{}
	err := db.Select(&messHalls, selectAllMessHallInfoQuery)
	if err != nil {
		return nil, err
	}

	if len(messHalls) == 0 {
		return nil, nil
	}

	return messHalls, nil
}

// GetAllMessHallAdminsInfoByID
func (pg *PostgresManager) GetMessHallAdminsInfoByID(id string) ([]usecases.MessHallAdmin, error) {

	db := pg.conn

	messHallAdmins := []usecases.MessHallAdmin{}
	query := fmt.Sprintf(selectMessHallAdminInfoByIDQuery, id)
	err := db.Select(&messHallAdmins, query)
	if err != nil {
		return nil, err
	}

	return messHallAdmins, nil
}

// GetAllMessHallAdminsInfo
func (pg *PostgresManager) GetAllMessHallAdminsInfo() ([]usecases.MessHallAdmin, error) {

	db := pg.conn

	messHallAdmins := []usecases.MessHallAdmin{}
	err := db.Select(&messHallAdmins, selectAllMessHallAdminInfoQuery)
	if err != nil {
		return nil, err
	}

	if len(messHallAdmins) == 0 {
		return nil, nil
	}

	return messHallAdmins, nil
}

// UpdateMessHallStatus
func (pg *PostgresManager) UpdateMessHallStatus(id string, status string) error {
	db := pg.conn

	_, err := db.Exec(updateMessHallStatusQuery, id, status)
	return err
}

// DeleteMessHall
func (pg *PostgresManager) DeleteMessHall(id string) error {
	db := pg.conn

	_, err := db.Exec(deleteMessHallQuery, id)
	_, err = db.Exec(deleteMessHallAdminQuery, id)
	return err
}

// AddMessHall creates a new mess hall entry
func (pg *PostgresManager) AddMessHall(messHall *usecases.MessHall, messHallAdmin *usecases.MessHallAdmin) error {

	db := pg.conn
	masshassTX := db.MustBegin()
	masshassadminsTX := db.MustBegin()

	//insert mess Hall info into table
	structTags := getListDBMEssHallTags(messHall)
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", "messhall", createQueryFields(structTags), createQueryValues(structTags))
	masshassTX.NamedExec(query, &messHall)
	err := masshassTX.Commit()
	if err != nil {
		log.Println(err)
		return err
	}

	//insert mess hall admin info into table
	structTags = getListDBMEssHallAdminTags(messHallAdmin)
	query = fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", "messhalls_admins", createQueryFields(structTags), createQueryValues(structTags))
	masshassadminsTX.NamedExec(query, &messHallAdmin)
	err = masshassadminsTX.Commit()
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

// GetMessHallMenuInfo
func (pg *PostgresManager) GetMessHallMenuInfo(messHallID string) ([]usecases.Menu, error) {

	db := pg.conn

	/* Get the messhall with this ID*/
	messHall, err := pg.GetMessHallsInfoByID(messHallID)
	if err != nil {
		return nil, err
	}

	if len(messHall) == 0 {
		return nil, nil
	}

	/* Get the menu for this mess hall */
	messHallMenu := []usecases.Menu{}
	query := fmt.Sprintf(selectMessHallMenuInfoQuery, messHall[0].MenuUID)
	err = db.Select(&messHallMenu, query)
	if err != nil {
		return nil, err
	}

	return messHallMenu, nil
}

// GetRecipesByRecipeUID
func (pg *PostgresManager) GetRecipesByRecipeUID(recipeUID string) []usecases.Recipe {

	db := pg.conn
	query := fmt.Sprintf(getRecipesByIDQuery, "recipe", recipeUID)

	recipes := []usecases.Recipe{}
	err := db.Select(&recipes, query)

	if err != nil {
		log.Println(err)
	}

	if len(recipes) == 0 {
		log.Println("Recipes table is empty")
	}

	return recipes
}

// IncrementMessHallAttendance
func (pg *PostgresManager) IncrementMessHallAttendance(id string) error {
	db := pg.conn

	_, err := db.Exec(incrementMessHallAttendanceQuery, id)
	return err
}

// DecrementMessHallAttendance
func (pg *PostgresManager) DecrementMessHallAttendance(id string) error {
	db := pg.conn

	_, err := db.Exec(decrementMessHallAttendanceQuery, id)
	return err
}
