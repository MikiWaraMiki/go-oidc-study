package auth

import (
	"context"
	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/oauth2"
	"log"
	"os"
)

type Authenticator struct {
	Provider *oidc.Provider
	Config oauth2.Config
	Ctx context.Context
}

func NewAuthenticator() (*Authenticator, error) {
	ctx := context.Background()
	issuer := "https://" + os.Getenv("AUTH0_DOMAIN") + "/"
	provider, err := oidc.NewProvider(ctx, issuer)

	if err != nil {
		log.Fatalf("failed provider intialize %v", err)
		return nil, err
	}

	config := oauth2.Config{
		ClientID: os.Getenv("AUTH0_CLIENT_ID"),
		ClientSecret: os.Getenv("AUTH0_CLIENT_SECRET"),
		RedirectURL: os.Getenv("AUTH0_CALLBACK_URL"),
		Endpoint: provider.Endpoint(),
		Scopes: []string{oidc.ScopeOpenID, "profile"},
	}

	return &Authenticator{
		provider,
		config,
		ctx,
	}, nil
}