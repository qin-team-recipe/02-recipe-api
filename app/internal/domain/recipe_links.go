package domain

type RecipeLinks struct {
	ID        int    `json:"id"`
	RecipeID  int    `json:"recipe_id"`
	Url       string `json:"url" binding:"required,http_url"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

type RecipeLinksForGet struct {
	ID       int    `json:"id"`
	RecipeID int    `json:"recipe_id"`
	Url      string `json:"url"`
}

func (rl *RecipeLinks) BuildForGet() *RecipeLinksForGet {
	return &RecipeLinksForGet{
		ID:       rl.ID,
		RecipeID: rl.RecipeID,
		Url:      rl.Url,
	}
}
