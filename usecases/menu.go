package usecases

// Menu
type Menu struct {
	MenuUID   string `db:"menu_uid" json:"menuUID"`
	RecipeUID string `db:"recipe_uid" json:"recipeUID"`
	TimeStamp []byte `db:"time_stamp" json:"timeStamp"`
}
