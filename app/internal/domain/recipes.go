package domain

type Recipes struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description *string `json:"description"`
	IsDraft     bool    `json:"is_draft"`
	CreatedAt   int64   `json:"created_at"`
	UpdatedAt   int64   `json:"updated_at"`
	DeletedAt   *int64  `json:"deleted_at"`
}

type RecipesForGet struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description *string `json:"description"`
	IsDraft     bool    `json:"is_draft"`
}

func (r *Recipes) BuildForGet() *RecipesForGet {
	return &RecipesForGet{
		ID:          r.ID,
		Title:       r.Title,
		Description: r.Description,
		IsDraft:     r.IsDraft,
	}
}
