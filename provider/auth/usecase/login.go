package usecase

import (
	"context"
	b64 "encoding/base64"
	"github.com/hinha/api-box/entity"
	"golang.org/x/oauth2"
)

type Login struct {
	State string
}

func (l *Login) GetCredentials(ctx context.Context, conf *oauth2.Config) (entity.LoginOAuthResponse, *entity.ApplicationError) {
	uEnc := b64.URLEncoding.EncodeToString([]byte(conf.AuthCodeURL(l.State)))
	return entity.LoginOAuthResponse{Link: uEnc, State: l.State}, nil
}
