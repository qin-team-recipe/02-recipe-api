package domain

type Users struct {
	ID          int    `json:"id"`
	ScreenName  string `json:"screen_name"`
	DisplayName string `json:"display_name"`
}
