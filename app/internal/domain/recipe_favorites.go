package domain

type RecipeFavorites struct {
	ID        int   `json:"id"`
	UserID    int   `json:"user_id"`
	RecipeID  int   `json:"recipe_id"`
	CreatedAt int64 `json:"created_at"`
}

type RecipeFavoritesForGet struct {
	ID       int `json:"id"`
	UserID   int `json:"user_id"`
	RecipeID int `json:"recipe_id"`

	Recipe *RecipesForGet `json:"recipe,omitempty"`
}

func (r *RecipeFavorites) BuildForGet() *RecipeFavoritesForGet {
	return &RecipeFavoritesForGet{
		ID:       r.ID,
		UserID:   r.UserID,
		RecipeID: r.RecipeID,
	}
}
