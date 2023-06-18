package domain

type UserRecipes struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	RecipeID  int    `json:"recipe_id"`
	CreatedAt int64  `json:"created_at"`
	DeletedAt *int64 `json:"deleted_at"`
}

type UserRecipesForGet struct {
	ID       int `json:"id"`
	UserID   int `json:"user_id"`
	RecipeID int `json:"recipe_id"`

	Recipe *RecipesForGet `json:"recipe"`
}

func (ur *UserRecipes) BuildForGet() *UserRecipesForGet {
	return &UserRecipesForGet{
		ID:       ur.ID,
		UserID:   ur.UserID,
		RecipeID: ur.RecipeID,
	}
}
