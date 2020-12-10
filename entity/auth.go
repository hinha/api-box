package entity

type LoginOAuthResponse struct {
	Link  string `json:"link"`
	State string `json:"state"`
}

type CallbackOAuth struct {
	Code  string `json:"link" query:"code"`
	State string `json:"state" query:"state"`
}

// User is a retrieved and authenticated user.
type GoogleUser struct {
	ID            string `json:"id"`
	Sub           string `json:"sub"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	Password      string `json:"password"`
	FamilyName    string `json:"family_name"`
	Profile       string `json:"profile"`
	Picture       string `json:"picture"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Gender        string `json:"gender"`
}
