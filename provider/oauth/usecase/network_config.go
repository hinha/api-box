package usecase

type GoogleOAuthNetworkConfig struct {
	Cfg struct {
		ClientID     string
		ClientSecret string
		RedirectURL  string
		Scopes       []string
		Endpoint     interface{}
	}
}

func (n *GoogleOAuthNetworkConfig) ClientID() string {
	return n.Cfg.ClientID
}

func (n *GoogleOAuthNetworkConfig) ClientSecret() string {
	return n.Cfg.ClientSecret
}

func (n *GoogleOAuthNetworkConfig) RedirectURL() string {
	return n.Cfg.RedirectURL
}

func (n *GoogleOAuthNetworkConfig) Scopes() []string {
	return n.Cfg.Scopes
}

func (n *GoogleOAuthNetworkConfig) Endpoint() interface{} {
	return n.Cfg.Endpoint
}
