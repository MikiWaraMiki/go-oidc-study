package handler

import (
	"context"
	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	"log"
	"net/http"
	"os"
)

func Index(c echo.Context) error {
	return c.JSON(http.StatusOK, "hello")
}

func Auth(c echo.Context) error {
	ctx := context.Background()
	issuer := "https//" + os.Getenv("AUTH0_DOMAIN") + "/"
	provider, err := oidc.NewProvider(ctx, issuer)

	if err != nil {
		log.Fatal(err)
	}

	config := oauth2.Config{
		ClientID: os.Getenv("AUTH0_CLIENT_ID"),
		ClientSecret: os.Getenv("AUTH0_CLIENT_SECRET"),
		RedirectURL: os.Getenv("/callback"),
		Endpoint: provider.Endpoint(),
		Scopes: []string{oidc.ScopeOpenID, "profile"},
	}

	state := "xyz" // 仮実装

	authURL := config.AuthCodeURL(state)

	return c.Redirect(http.StatusMovedPermanently, authURL)
}