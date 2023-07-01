package domain

type Recipes struct {
	ID          int     `json:"id"`
	Title       string  `json:"title" binding:"required"`
	Description *string `json:"description"`
	Servings    int     `json:"servings" binding:"required,min=1"`
	IsDraft     bool    `json:"is_draft"`
	CreatedAt   int64   `json:"created_at"`
	UpdatedAt   int64   `json:"updated_at"`
	DeletedAt   *int64  `json:"deleted_at"`
}

type RecipesForGet struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description *string `json:"description"`
	Servings    int     `json:"servings"`
	IsDraft     bool    `json:"is_draft"`

	FavoritesCount int `json:"facorites_count"`
}

func (r *Recipes) BuildForGet() *RecipesForGet {
	return &RecipesForGet{
		ID:          r.ID,
		Title:       r.Title,
		Description: r.Description,
		Servings:    r.Servings,
		IsDraft:     r.IsDraft,
	}
}
