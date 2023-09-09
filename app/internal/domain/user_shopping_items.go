package domain

type UserShoppingItems struct {
	ID          int     `json:"id"`
	UserID      int     `json:"user_id"`
	Title       string  `json:"title" binding:"required"`
	Description *string `json:"description"`
	IsDone      bool    `json:"is_done"`
	CreatedAt   int64   `json:"created_at"`
	UpdatedAt   int64   `json:"updated_at"`
}

type UserShoppingItemsForGet struct {
	ID          int     `json:"id"`
	UserID      int     `json:"user_id"`
	Title       string  `json:"title" binding:"required"`
	Description *string `json:"description"`
	IsDone      bool    `json:"is_done"`
}

func (u *UserShoppingItems) BuildForGet() *UserShoppingItemsForGet {
	return &UserShoppingItemsForGet{
		ID:          u.ID,
		UserID:      u.UserID,
		Title:       u.Title,
		Description: u.Description,
		IsDone:      u.IsDone,
	}
}
