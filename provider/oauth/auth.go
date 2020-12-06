package oauth

import (
	"context"
	"github.com/hinha/api-box/entity"
	"github.com/hinha/api-box/provider"
	"github.com/hinha/api-box/provider/oauth/api"
	"github.com/hinha/api-box/provider/oauth/usecase"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"os"
)

//type GoogleOAuth struct {
//	network provider.GoogleOAuth
//	googleNetworkCfg provider.GoogleOAuthNetworkConfig
//}
//
//func Fabricate(network provider.GoogleOAuth) *GoogleOAuth {
//	networkCfg := &usecase.GoogleOAuthNetworkConfig{}
//	networkCfg.Cfg.ClientID = os.Getenv("GoogleClientID")
//	networkCfg.Cfg.ClientSecret = os.Getenv("GoogleClientSecret")
//	networkCfg.Cfg.RedirectURL = os.Getenv("GoogleRedirectURL")
//	networkCfg.Cfg.Scopes = []string{"https://www.googleapis.com/auth/userinfo.email"}
//	return &GoogleOAuth{
//		network: network,
//		googleNetworkCfg: networkCfg,
//	}
//}

type GoogleOAuth struct {
	conf *oauth2.Config
}

func FabricateGoogle() *GoogleOAuth {
	conf := &oauth2.Config{}
	conf.ClientID = os.Getenv("GoogleClientID")
	conf.ClientSecret = os.Getenv("GoogleClientSecret")
	conf.RedirectURL = os.Getenv("GoogleRedirectURL")
	conf.Scopes = []string{"https://www.googleapis.com/auth/userinfo.profile"}
	conf.Endpoint = google.Endpoint
	return &GoogleOAuth{conf: conf}
}

func (g *GoogleOAuth) FabricateAPI(engine provider.APIEngine) {
	engine.InjectAPI(api.NewLogin(g))
	engine.InjectAPI(api.NewCallbackLogin(g))
}

func (g *GoogleOAuth) CallbackURIGoogle(ctx context.Context, request entity.CallbackOAuth) (entity.GoogleUser, *entity.ApplicationError) {
	codeCallback := &usecase.Code{}
	return codeCallback.ParseCode(ctx, request, g.conf)
}

// Login ... get prefer link
func (g *GoogleOAuth) Login(ctx context.Context, state string) (entity.LoginOAuthResponse, *entity.ApplicationError) {
	login := &usecase.Login{State: state}
	return login.GetCredentials(ctx, g.conf)
}
