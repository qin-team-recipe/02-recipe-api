package domain

type ChefRecipes struct {
	ID        int    `json:"id"`
	ChefID    int    `json:"chef_id"`
	RecipeID  int    `json:"recipe_id"`
	CreatedAt int64  `json:"created_at"`
	DeletedAt *int64 `json:"deleted_at"`
}

type ChefRecipesForGet struct {
	ID       int `json:"id"`
	ChefID   int `json:"chef_id"`
	RecipeID int `json:"recipe_id"`

	Recipe *RecipesForGet `json:"recipe"`
}

func (cr *ChefRecipes) BuildForGet() *ChefRecipesForGet {
	return &ChefRecipesForGet{
		ID:       cr.ID,
		ChefID:   cr.ChefID,
		RecipeID: cr.RecipeID,
	}
}
