package api

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/sessions"
	"github.com/hinha/api-box/provider"
	"github.com/hinha/api-box/provider/api/command"
	"github.com/hinha/api-box/provider/api/handler"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"os"
	"strconv"
)

// API ...
type API struct {
	engine *echo.Echo
	port   int
}

func Fabricate(givenPort int) *API {
	return &API{
		engine: echo.New(),
		port:   givenPort,
	}
}

// FabricateCommand insert api related command
func (a *API) FabricateCommand(cmd provider.Command) {
	cmd.InjectCommand(
		command.NewRun(a),
	)
}

// InjectAPI inject new API into api_box
func (a *API) InjectAPI(handler provider.APIHandler) {
	a.engine.Add(handler.Method(), handler.Path(), func(context echo.Context) error {
		req := context.Request()
		if reqID := req.Header.Get("X-Request-ID"); reqID != "" {
			context.Set("request-id", reqID)
		} else {
			context.Set("request-id", uuid.New().String())
		}

		if userID := req.Header.Get("Resource-Owner-ID"); userID != "" {
			convertedUserID, err := strconv.Atoi(userID)
			if err == nil {
				context.Set("user-id", convertedUserID)
			}
		}
		sess, _ := session.Get("session", context)
		sess.Options = &sessions.Options{
			Path:     "/",
			MaxAge:   86400 * 7,
			HttpOnly: true,
		}

		handler.Handle(context, sess)
		return nil
	})
}

func (a *API) Run() error {
	a.engine.Use(
		middleware.Logger(),
		session.Middleware(sessions.NewCookieStore([]byte(os.Getenv("APP_ENCRYPTION_KEY")))),
	)
	a.InjectAPI(handler.NewHealth())
	return a.engine.Start(fmt.Sprintf(":%d", a.port))
}

// Shutdown api engine
func (a *API) Shutdown(ctx context.Context) error {
	return a.engine.Shutdown(ctx)
}
