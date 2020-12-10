package auth

import (
	"context"
	"github.com/hinha/api-box/entity"
	"github.com/hinha/api-box/provider"
	"github.com/hinha/api-box/provider/auth/api"
	"github.com/hinha/api-box/provider/auth/usecase"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"os"
)

type GoogleOAuth struct {
	conf         *oauth2.Config
	userProvider provider.UserOAuth
}

func FabricateGoogle(userProvider provider.UserOAuth) *GoogleOAuth {
	conf := &oauth2.Config{}
	conf.ClientID = os.Getenv("GoogleClientID")
	conf.ClientSecret = os.Getenv("GoogleClientSecret")
	conf.RedirectURL = os.Getenv("GoogleRedirectURL")
	conf.Scopes = []string{
		"https://www.googleapis.com/auth/userinfo.email",
		"https://www.googleapis.com/auth/userinfo.profile",
	}
	conf.Endpoint = google.Endpoint
	return &GoogleOAuth{conf: conf, userProvider: userProvider}
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
