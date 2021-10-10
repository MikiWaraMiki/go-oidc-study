package main

import (
	"github.com/MikiWaraMiki/go-oidc-study/handler"
	"github.com/labstack/echo"
	"github.com/labstack/echo/v4/middleware"
)

func newRouter() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", handler.Index)

	return e
}