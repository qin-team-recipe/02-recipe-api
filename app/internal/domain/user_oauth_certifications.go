package domain

type UserOauthCertifications struct {
	ID            int    `json:"id"`
	UserID        int    `json:"user_id"`
	ServiceUserID string `json:"service_user_id"`
	ServiceName   string `json:"service_name"`
	CreatedAt     int64  `json:"created_at"`
}
