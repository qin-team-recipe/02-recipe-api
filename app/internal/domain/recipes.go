package domain

type Recipes struct {
	ID          int     `json:"id"`
	ChefID      int     `json:"chef_id"`
	Title       string  `json:"title"`
	Description *string `json:"description"`
	CreatedAt   int64   `json:"created_at"`
	UpdatedAt   int64   `json:"updated_at"`
	DeletedAt   *int64  `json:"deleted_at"`
}

type RecipesForGet struct {
	ID          int     `json:"id"`
	ChefID      int     `json:"chef_id"`
	Title       string  `json:"title"`
	Description *string `json:"description"`
}

func (r *Recipes) BuildForGet() *RecipesForGet {
	return &RecipesForGet{
		ID:          r.ID,
		ChefID:      r.ChefID,
		Title:       r.Title,
		Description: r.Description,
	}
}
