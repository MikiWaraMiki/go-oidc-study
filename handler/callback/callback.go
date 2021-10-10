package callback

import (
	"context"
	"encoding/gob"
	"fmt"
	"github.com/MikiWaraMiki/go-oidc-study/auth"
	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"os"
)

type ResponseUserDto struct {
	Name string `json:"name"`
	FamilyName string `json:"familyName"`
	GivenName string `json:"givenName"`
}

func Handler(c echo.Context) error {
	log.Printf("callback started")
	// Get State from session
	sess, err := session.Get("sample", c)
	if err != nil {
		log.Fatalf("failed session %v", err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	state := sess.Values["state"]
	// Verify State
	if c.QueryParam("state") != state {
		log.Printf("invalid state")
		return c.JSON(http.StatusBadRequest, "invalid state")
	}
	// Generate authenticator
	authenticator, err := auth.NewAuthenticator()
	if err != nil {
		log.Fatalf("failed gen authenticator %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	// Exchange id token
	log.Printf("%v", c.QueryParam("state"))
	token, err := authenticator.Config.Exchange(context.TODO(), c.QueryParam("code"))
	if err != nil {
		log.Printf("unauthorized")
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	}
	rawIdToken, ok := token.Extra("id_token").(string)
	if !ok {
		log.Fatalf("failed fetch id_token ")
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	// verify id token
	oidcConfig := &oidc.Config{
		ClientID: os.Getenv("AUTH0_CLIENT_ID"),
	}
	idToken, err := authenticator.Provider.Verifier(oidcConfig).Verify(context.TODO(), rawIdToken)
	if err != nil {
		log.Fatalf("failed to verify token: %v", err)
		return err
	}

	// get profile data
	var profile map[string]interface{}

	if err := idToken.Claims(&profile); err != nil {
		log.Fatalf("failed get user profile data %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	fmt.Printf("%#v", profile)
	gob.Register(map[string]interface{}{})
	sess.Values["id_token"] = rawIdToken
	sess.Values["access_token"] = token.AccessToken
	sess.Values["profile"] = profile
	err = sess.Save(c.Request(), c.Response())
	if err != nil {
		log.Fatalf("failed save session %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	if err != nil {
		log.Fatalf("Failed convert json", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, profile)
}