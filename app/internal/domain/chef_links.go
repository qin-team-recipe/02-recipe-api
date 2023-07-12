package domain

type ChefLinks struct {
	ID          int     `json:"id"`
	ChefID      int     `json:"chef_id"`
	ServiceName *string `json:"service_name"`
	Url         string  `json:"url" binding:"required,http_url"`
	CreatedAt   int64   `json:"created_at"`
	UpdatedAt   int64   `json:"updated_at"`
}

type ChefLinksForGet struct {
	ID          int     `json:"id"`
	ChefID      int     `json:"chef_id"`
	ServiceName *string `json:"service_name"`
	Url         string  `json:"url"`
}

func (c *ChefLinks) BuildForGet() *ChefLinksForGet {
	return &ChefLinksForGet{
		ID:          c.ID,
		ChefID:      c.ChefID,
		ServiceName: c.ServiceName,
		Url:         c.Url,
	}
}
