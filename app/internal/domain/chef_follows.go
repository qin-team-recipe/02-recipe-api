package domain

type ChefFollows struct {
	ID        int   `json:"id"`
	UserID    int   `json:"user_id"`
	ChefID    int   `json:"chef_id"`
	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`
}

type ChefFollowsForGet struct {
	ID     int `json:"id"`
	UserID int `json:"user_id"`
	ChefID int `json:"chef_id"`

	Chef *ChefsForGet `json:"chef"`
}

func (c *ChefFollows) BuildForGet() *ChefFollowsForGet {
	return &ChefFollowsForGet{
		ID:     c.ID,
		UserID: c.UserID,
		ChefID: c.ChefID,
	}
}
