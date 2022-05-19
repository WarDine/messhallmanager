package repositories

import (
	// "recipemanager/usecases"
	// "database/sql"
	// "encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
	"usecases"

	// "recipemanager/usecases"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PostgresManager struct {
	conn *sqlx.DB
}

var PostgresRepo PostgresManager

// return postgres objects which contains connection to database
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

func (pg *PostgresManager) GetAllMessHallsInfo() ([]usecases.MessHall, error) {

	db := pg.conn

	messHalls := []usecases.MessHall{}
	err := db.Select(&messHalls, "SELECT * FROM mass_hall;")
	if err != nil {
		return nil, err
	}

	if len(messHalls) == 0 {
		return nil, errors.New("MessHall table is empty")
	}

	return messHalls, nil
}

func GetListDBMEssHallTags(messHall *usecases.MessHall) []string {

	t := reflect.TypeOf(*messHall)

	tagFields := make([]string, t.NumField())
	for i := range tagFields {
		tagFields[i] = GetDBTagName(messHall, t.Field(i).Name)
	}

	return tagFields
}

func (pg *PostgresManager) AddMessHall(messHall *usecases.MessHall) error {

	db := pg.conn
	tx := db.MustBegin()

	structTags := GetListDBMEssHallTags(messHall)
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", "mass_hall", createQueryFields(structTags), createQueryValues(structTags))

	tx.NamedExec(query, &messHall)
	err := tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (pg *PostgresManager) PingConnection() {

	err := pg.conn.Ping()
	if err != nil {
		panic(err)
	} else {
		log.Println("Connection works as expected")
	}
}

func (pg *PostgresManager) DeleteConnection() {
	defer pg.conn.Close()
}

type Ingredient struct {
	Ingredient_uid int    `db:"ingredient_uid" json:"ingredientUID"`
	Name           string `db:"ingredient_name" json:"name"`
	RecipeUID      int    `db:"recipe_uid" json:"recipeUID"`
	Amount         int    `db:"amount" json:"amount"`
}

// get list of all db tags of a struct ingredient
// call with GetListDBTags(ingredient)
func GetListDBTags(genericStruct *Ingredient) []string {

	t := reflect.TypeOf(*genericStruct)

	tagFields := make([]string, t.NumField())
	for i := range tagFields {
		tagFields[i] = GetDBTagName(genericStruct, t.Field(i).Name)
	}

	return tagFields
}

// get db tag name of a field from a generic struct
func GetDBTagName(genericStruct interface{}, structField string) string {

	tagName := "db"
	field, ok := reflect.TypeOf(genericStruct).Elem().FieldByName(structField)
	if !ok {
		log.Fatal("Field not found")
	}

	return string(field.Tag.Get(tagName))
}

/**
 * this generates a string like:
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
 * this generates a string like:
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

// example of query
// tx.NamedExec("INSERT INTO ingredient (ingredient_name, recipe_uid, amount) VALUES (:ingredient_name, :recipe_uid, :amount)", &second_ingredient)
func (pg *PostgresManager) InsertIngredientIntoDB(tableName string, ingredient *Ingredient) {

	structTags := GetListDBTags(ingredient)
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", tableName, createQueryFields(structTags), createQueryValues(structTags))
	db := pg.conn

	tx := db.MustBegin()

	tx.NamedExec(query, &ingredient)
	err := tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
}

func (pg *PostgresManager) GetAllIngredients(tableName string) []Ingredient {

	db := pg.conn

	ingredients := []Ingredient{}
	err := db.Select(&ingredients, "SELECT * FROM ingredient;")
	if err != nil {
		log.Fatal(err)
	}

	if len(ingredients) == 0 {
		log.Println("Ingredients table is empty")
	}

	return ingredients
}

func convertfilterToString(m map[string]interface{}) {

}

/**
 * filter is a map[string]interface{}
 * it will be converted into a string an added in query
 */
// WIP
func (pg *PostgresManager) GetFilteredIngredients(tableName string, filter string) []Ingredient {

	db := pg.conn

	if filter == "" {
		log.Print("Filter is empty; Getting all values:")
		return pg.GetAllIngredients(tableName)
	}

	ingredients := []Ingredient{}
	err := db.Select(&ingredients, "SELECT * FROM ingredient;")
	if err != nil {
		log.Fatal(err)
	}

	if len(ingredients) == 0 {
		log.Println("Ingredients table is empty")
	}

	return ingredients
}

// this function is just an example of how to use sqlx lib
func (pg *PostgresManager) TestDatabase(tableName string) {

	db := pg.conn

	tx := db.MustBegin()
	var secondIngredient = &Ingredient{
		Name:      "Pastarnac",
		RecipeUID: 1918273,
		Amount:    2,
	}

	tx.MustExec("INSERT INTO ingredient (ingredient_name, recipe_uid, amount) VALUES ($1, $2, $3)", "Telina", 1918273, 2)
	tx.NamedExec("INSERT INTO ingredient (ingredient_name, recipe_uid, amount) VALUES (:ingredient_name, :recipe_uid, :amount)", &secondIngredient)
	tx.MustExec("INSERT INTO ingredient (ingredient_name, recipe_uid, amount) VALUES ($1, $2, $3)", "Morcov", 1918273, 4)
	tx.Commit()

	ingredients := []Ingredient{}

	err := db.Select(&ingredients, "SELECT * FROM ingredient;")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Len of ingredients: %d\n", len(ingredients))
	if len(ingredients) >= 2 {
		ingredient1, ingredient2 := ingredients[0], ingredients[1]
		log.Printf("%#v\n%#v", ingredient1, ingredient2)
	}

	pastarnac := Ingredient{}
	err = db.Get(&pastarnac, "SELECT * FROM ingredient WHERE ingredient_name=$1;", "Pastarnac")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Pastarnacul: %#v\n", pastarnac)

	pg.DeleteConnection()
}
