package domain

type GoogleUserAccount struct {
	GoogleUserID string `json:"google_user_id"`
	DisplayName  string `json:"display_name"`
	Email        string `json:"email"`
}
