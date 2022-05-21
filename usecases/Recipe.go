package usecases

type Recipe struct {
	RecipeUID    string `db:"recipe_uid" json:"recipeUID,omitempty"`
	MesshallUID  string `db:"messhall_uid" json:"messhallUID,omitempty"`
	Name         string `db:"nname" json:"name"`
	Description  string `db:"description" json:"description"`
	Calories     int    `db:"calories" json:"calories"`
	CookingTime  int    `db:"cooking_time" json:"cookingTime"`
	Instructions string `db:"instructions" json:"instructions,omitempty"`
	Portions     int    `db:"portions" json:"portions,omitempty"`
}
