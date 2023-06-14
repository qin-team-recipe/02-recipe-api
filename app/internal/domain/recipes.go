package domain

type Recipes struct {
	ID           int     `json:"id"`
	ChefRecipeID int     `json:"chef_recipe_id"`
	Title        string  `json:"title"`
	Description  *string `json:"description"`
	CreatedAt    int64   `json:"created_at"`
	UpdatedAt    int64   `json:"updated_at"`
	DeletedAt    *int64  `json:"deleted_at"`
}

type RecipesForGet struct {
	ID           int     `json:"id"`
	ChefRecipeID int     `json:"chef_recipe_id"`
	Title        string  `json:"title"`
	Description  *string `json:"description"`
}

func (r *Recipes) BuildForGet() *RecipesForGet {
	return &RecipesForGet{
		ID:           r.ID,
		ChefRecipeID: r.ChefRecipeID,
		Title:        r.Title,
		Description:  r.Description,
	}
}
