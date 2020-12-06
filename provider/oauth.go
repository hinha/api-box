package provider

import (
	"context"
	"github.com/hinha/api-box/entity"
)

type OAuth interface {
	Login(ctx context.Context, stateKey string) (entity.LoginOAuthResponse, *entity.ApplicationError)
	CallbackURIGoogle(ctx context.Context, request entity.CallbackOAuth) (entity.GoogleUser, *entity.ApplicationError)
}
