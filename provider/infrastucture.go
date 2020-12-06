package provider

import (
	"context"
)

type GoogleOAuth interface {
	GET(ctx context.Context, cfg GoogleOAuthNetworkConfig, path string)
}

type GoogleOAuthNetworkConfig interface {
	ClientID() string
	ClientSecret() string
	RedirectURL() string
	Scopes() []string
	Endpoint() interface{}
}

//type Network interface {
//	GET(ctx context.Context, cfg NetworkConfig, oauthCfg *oauth2.Config, path string)
//}
//
// NetworkConfig given for network request
//type NetworkConfig interface {
//	Host() string
//	Username() string
//	Password() string
//	Timeout() time.Duration
//	Retry() int
//	RetrySleepDuration() time.Duration
//}
