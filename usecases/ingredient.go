package usecases

type Ingredient struct {
	Ingredient_uid int    `db:"ingredient_uid" json:"ingredientUID"`
	Name           string `db:"ingredient_name" json:"name"`
	RecipeUID      int    `db:"recipe_uid" json:"recipeUID"`
	Amount         int    `db:"amount" json:"amount"`
}
