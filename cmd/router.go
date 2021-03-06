package cmd

import (
	"github.com/MikiWaraMiki/go-oidc-study/handler"
	"github.com/MikiWaraMiki/go-oidc-study/handler/callback"
	"github.com/MikiWaraMiki/go-oidc-study/handler/login"
	"github.com/MikiWaraMiki/go-oidc-study/handler/logout"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func newRouter() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("sample"))))

	e.GET("/", handler.Index)
	e.GET("/auth", login.Handler)
	e.GET("/callback", callback.Handler)
	e.GET("/logout", logout.Handler)

	return e
}
