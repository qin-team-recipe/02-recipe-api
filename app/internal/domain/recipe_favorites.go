package domain

type RecipeFavorites struct {
	ID        int   `json:"id"`
	UserID    int   `json:"user_id"`
	RecipeID  int   `json:"recipe_id"`
	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`
}
