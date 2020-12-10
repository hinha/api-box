package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/sessions"
	"github.com/hinha/api-box/entity"
	"github.com/hinha/api-box/provider"
	"net/http"
)

type CallbackURI struct {
	oauthProvider provider.OAuth
}

// NewLogin create new request otp handler object
func NewCallbackLogin(authProvider provider.OAuth) *CallbackURI {
	return &CallbackURI{oauthProvider: authProvider}
}

// Path return api path
func (r *CallbackURI) Path() string {
	return "/authorization/:provider/callback"
}

// Method return api method
func (r *CallbackURI) Method() string {
	return "GET"
}

func (r *CallbackURI) Handle(context provider.APIContext, sess *sessions.Session) {
	var request entity.CallbackOAuth
	code := context.QueryParam("code")
	if code == "" {
		_ = context.JSON(http.StatusBadRequest, map[string]interface{}{
			"errors":  []string{"bad request given by client"},
			"message": "Bad request",
		})
		return
	}
	request.Code = code

	switch context.Param("provider") {
	case "google":
		users, err := r.oauthProvider.CallbackURIGoogle(context.Request().Context(), request)
		if err != nil {
			_ = context.JSON(err.HTTPStatus, map[string]interface{}{
				"errors":  err.ErrorString(),
				"message": err.Error(),
			})
			return
		}

		jsonString, _ := json.Marshal(users)
		sess.Values["google-user"] = string(jsonString)
		_ = sess.Save(context.Request(), context.Response())
		_ = context.JSON(http.StatusOK, map[string]interface{}{
			"data": users,
		})
	default:
		_ = context.JSON(http.StatusBadRequest, map[string]interface{}{
			"errors":  []string{fmt.Sprintf("provider not found %s", context.Param("provider"))},
			"message": "Bad request",
		})
		return
	}
}
