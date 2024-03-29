package domain

type Recipes struct {
	ID              int     `json:"id"`
	WatchID         string  `json:"watch_id"`
	Title           string  `json:"title" binding:"required"`
	Description     *string `json:"description"`
	Servings        int     `json:"servings" binding:"required,min=1"`
	IsDraft         bool    `json:"is_draft"`
	PublishedStatus string  `json:"published_status"`
	CreatedAt       int64   `json:"created_at"`
	UpdatedAt       int64   `json:"updated_at"`
	DeletedAt       *int64  `json:"deleted_at"`
}

type RecipesForGet struct {
	ID              int     `json:"id"`
	WatchID         string  `json:"watch_id"`
	Title           string  `json:"title"`
	Description     *string `json:"description"`
	Servings        int     `json:"servings"`
	IsDraft         bool    `json:"is_draft"`
	PublishedStatus string  `json:"published_status"`

	FavoritesCount int `json:"favorites_count"`

	// この値はどちらかが入る
	Chef *ChefsForGet `json:"chef,omitempty"`
	User *UsersForGet `json:"user,omitempty"`
}

func (r *Recipes) BuildForGet() *RecipesForGet {
	return &RecipesForGet{
		ID:              r.ID,
		WatchID:         r.WatchID,
		Title:           r.Title,
		Description:     r.Description,
		Servings:        r.Servings,
		IsDraft:         r.IsDraft,
		PublishedStatus: r.PublishedStatus,
	}
}
