package api

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gorilla/sessions"
	"github.com/hinha/api-box/entity"
	"github.com/hinha/api-box/provider"
	"net/http"
)

type Login struct {
	oauthProvider provider.OAuth
}

// NewLogin create new request otp handler object
func NewLogin(authProvider provider.OAuth) *Login {
	return &Login{oauthProvider: authProvider}
}

// Path return api path
func (r *Login) Path() string {
	return "/authorization/:provider/login"
}

// Method return api method
func (r *Login) Method() string {
	return "GET"
}

// Handle request google auth
func (r *Login) Handle(context provider.APIContext, sess *sessions.Session) {
	state, err := r.RandToken(32)
	if err != nil {
		_ = context.JSON(http.StatusBadRequest, map[string]interface{}{
			"errors":  []string{"bad request given by client", "Error while generating random data."},
			"message": "Bad request",
		})
		return
	}

	if val, ok := sess.Values["google-user"]; ok {
		var users entity.GoogleUser
		if val != "" {
			strMarshal := fmt.Sprintf("%v", val)
			json.Unmarshal([]byte(strMarshal), &users)

			_ = context.JSON(http.StatusOK, map[string]interface{}{
				"data": users,
			})
			return
		}
	}

	switch context.Param("provider") {
	case "google":
		sess.Values["state"] = state
		_ = sess.Save(context.Request(), context.Response())

		response, err := r.oauthProvider.Login(context.Request().Context(), state)
		if err != nil {
			_ = context.JSON(err.HTTPStatus, map[string]interface{}{
				"errors":  err.ErrorString(),
				"message": err.Error(),
			})
			return
		}

		_ = context.JSON(http.StatusOK, map[string]interface{}{
			"data": response,
		})
	default:
		_ = context.JSON(http.StatusBadRequest, map[string]interface{}{
			"errors":  []string{fmt.Sprintf("provider not found %s", context.Param("provider"))},
			"message": "Bad request",
		})
		return
	}
}

// RandToken generates a random @l length token.
func (r *Login) RandToken(l int) (string, error) {
	b := make([]byte, l)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(b), nil
}
