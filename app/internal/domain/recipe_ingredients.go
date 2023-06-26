package domain

type RecipeIngredients struct {
	ID          int     `json:"id"`
	RecipeID    int     `json:"recipe_id" binding:"required"`
	Name        string  `json:"name" binding:"required"`
	Description *string `json:"description"`
	CreatedAt   int64   `json:"created_at"`
	UpdatedAt   int64   `json:"updated_at"`
}

type RecipeIngredientsForGet struct {
	ID          int     `json:"id"`
	RecipeID    int     `json:"recipe_id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

func (r *RecipeIngredients) BuildForGet() *RecipeIngredientsForGet {
	return &RecipeIngredientsForGet{
		ID:          r.ID,
		RecipeID:    r.RecipeID,
		Name:        r.Name,
		Description: r.Description,
	}
}
