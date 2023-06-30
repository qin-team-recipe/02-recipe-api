package domain

type SocialUserAccount struct {
	ServiceName   string `json:"service_name"`
	ServiceUserID string `json:"service_user_id"`
	DisplayName   string `json:"display_name"`
	Email         string `json:"email"`
}

type SocialServiceType struct {
	Google    string
	Instagram string
}

func NewSocialServiceType() *SocialServiceType {
	return &SocialServiceType{
		Google:    "google",
		Instagram: "instagram",
	}
}
