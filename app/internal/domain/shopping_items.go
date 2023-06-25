package domain

type ShoppingItems struct {
	ID                 int   `json:"id"`
	UserID             int   `json:"user_id" binding:"required"`
	RecipeIngredientID int   `json:"recipe_ingredient_id" binding:"required"`
	IsDone             bool  `json:"is_done"`
	CreatedAt          int64 `json:"created_at"`
	UpdatedAt          int64 `json:"updated_at"`
}

type ShoppingItemsForGet struct {
	ID                 int  `json:"id"`
	UserID             int  `json:"user_id"`
	RecipeIngredientID int  `json:"recipe_ingredient_id"`
	IsDone             bool `json:"is_done"`

	RecipeIngredient *RecipeIngredientsForGet `json:"recipe_ingredient"`
}

func (s *ShoppingItems) BuildForGet() *ShoppingItemsForGet {
	return &ShoppingItemsForGet{
		ID:                 s.ID,
		UserID:             s.UserID,
		RecipeIngredientID: s.RecipeIngredientID,
		IsDone:             s.IsDone,
	}
}
