package usecase

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/hinha/api-box/entity"
	"github.com/hinha/api-box/provider"
	"golang.org/x/oauth2"
	"io/ioutil"
	"net/http"
)

type Code struct {
	State string
}

func (l *Code) ParseCode(ctx context.Context, request entity.CallbackOAuth, conf *oauth2.Config) (entity.GoogleUser, *entity.ApplicationError) {

	// get code from client ex: 4/0AY0e-xxx
	token, err := conf.Exchange(ctx, request.Code)
	if err != nil {
		return entity.GoogleUser{}, l.invalidAuthError("Login failed. credential invalid.")
	}

	// Verify token
	client := conf.Client(ctx, token)
	userinfo, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return entity.GoogleUser{}, l.invalidAuthError("Login failed. Please try again.")
	}
	defer userinfo.Body.Close()

	data, _ := ioutil.ReadAll(userinfo.Body)
	u := entity.GoogleUser{}
	if err = json.Unmarshal(data, &u); err != nil {
		return entity.GoogleUser{}, l.invalidAuthError("Error marshalling response. Please try again.")
	}

	return u, nil
}

func (l *Login) PerformOAuth(ctx context.Context, response entity.GoogleUser, userProvider provider.UserOAuth) *entity.ApplicationError {
	_, errProvider := userProvider.CreateOAuth(ctx, response)
	if errProvider != nil {
		return errProvider
	}

	return nil
}

func (l *Code) invalidParseURLError() *entity.ApplicationError {
	return &entity.ApplicationError{
		Err:        []error{errors.New("url not invalid")},
		HTTPStatus: http.StatusBadRequest,
	}
}

func (l *Code) invalidBase64Error() *entity.ApplicationError {
	return &entity.ApplicationError{
		Err:        []error{errors.New("payload bearer invalid")},
		HTTPStatus: http.StatusBadRequest,
	}
}

func (l *Code) invalidAuthError(message string) *entity.ApplicationError {
	return &entity.ApplicationError{
		Err:        []error{errors.New(message)},
		HTTPStatus: http.StatusBadRequest,
	}
}
