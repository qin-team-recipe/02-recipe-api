package domain

type Chefs struct {
	ID          int     `json:"id"`
	ScreenName  string  `json:"screen_name"`
	DisplayName string  `json:"display_name"`
	Description *string `json:"description"`
	CreatedAt   int64   `json:"created_at"`
	UpdatedAt   int64   `json:"updated_at"`
	DeletedAt   *int64  `json:"deleted_at"`
}

type ChefsForGet struct {
	ID          int     `json:"id"`
	ScreenName  string  `json:"screen_name"`
	DisplayName string  `json:"display_name"`
	Description *string `json:"description"`
}

func (c *Chefs) BuildForGet() *ChefsForGet {
	return &ChefsForGet{
		ID:          c.ID,
		ScreenName:  c.ScreenName,
		DisplayName: c.DisplayName,
		Description: c.Description,
	}
}
