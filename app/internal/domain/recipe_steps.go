package domain

type RecipeSteps struct {
	ID          int     `json:"id"`
	RecipeID    int     `json:"recipe_id"`
	Title       string  `json:"title" binding:"required"`
	Description *string `json:"description"`
	StepNumber  int     `json:"step_number"`
	CreatedAt   int64   `json:"created_at"`
	UpdatedAt   int64   `json:"updated_at"`
}

type RecipeStepsForGet struct {
	ID          int     `json:"id"`
	RecipeID    int     `json:"recipe_id"`
	Title       string  `json:"title"`
	Description *string `json:"description"`
	StepNumber  int     `json:"step_number"`
}

func (r *RecipeSteps) BuildForGet() *RecipeStepsForGet {
	return &RecipeStepsForGet{
		ID:          r.ID,
		RecipeID:    r.RecipeID,
		Title:       r.Title,
		Description: r.Description,
		StepNumber:  r.StepNumber,
	}
}
