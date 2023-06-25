package domain

type ShoppingMemos struct {
	ID                 int   `json:"id"`
	UserID             int   `json:"user_id" binding:"required"`
	RecipeIngredientID int   `json:"recipe_ingredient_id" binding:"required"`
	IsDone             bool  `json:"is_done"`
	CreatedAt          int64 `json:"created_at"`
	UpdatedAt          int64 `json:"updated_at"`
}

type ShoppingMemosForGet struct {
	ID                 int  `json:"id"`
	UserID             int  `json:"user_id"`
	RecipeIngredientID int  `json:"recipe_ingredient_id"`
	IsDone             bool `json:"is_done"`

	RecipeIngredient *RecipeIngredientsForGet `json:"recipe_ingredient"`
}

func (s *ShoppingMemos) BuildForGet() *ShoppingMemosForGet {
	return &ShoppingMemosForGet{
		ID:                 s.ID,
		UserID:             s.UserID,
		RecipeIngredientID: s.RecipeIngredientID,
		IsDone:             s.IsDone,
	}
}
