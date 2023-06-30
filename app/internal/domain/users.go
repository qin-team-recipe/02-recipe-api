package domain

type Users struct {
	ID          int    `json:"id"`
	ScreenName  string `json:"screen_name"`
	DisplayName string `json:"display_name"`
	Email       string `json:"email"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
	DeletedAt   int64  `json:"deleted_at"`
}

type UsersForGet struct {
	ID          int    `json:"id"`
	ScreenName  string `json:"screen_name"`
	DisplayName string `json:"display_name"`
	Email       string `json:"email"`
}

func (u *Users) BuildForGet() *UsersForGet {
	return &UsersForGet{
		ID:          u.ID,
		ScreenName:  u.ScreenName,
		DisplayName: u.DisplayName,
		Email:       u.Email,
	}
}
