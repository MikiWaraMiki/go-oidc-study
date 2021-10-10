package login

import (
	"encoding/base64"
	"github.com/MikiWaraMiki/go-oidc-study/auth"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"log"
	"math/rand"
	"net/http"
)

func LoginHandler(c echo.Context) error {
	// Gen random state
	b := make([]byte, 32)
	_, err := rand.Read(b)

	if err != nil {
		log.Fatalf("failed gen state value %v", err)
		return c.JSON(http.StatusInternalServerError, "failed state")
	}
	state := base64.StdEncoding.EncodeToString(b)

	// Save state value to Session
	sess, _ := session.Get("sample", c)
	sess.Options = &sessions.Options{
		Path: "/",
		MaxAge: 86400 * 7,
		HttpOnly: true,
	}
	sess.Values["state"] = state
	err = sess.Save(c.Request(), c.Response())
	if err != nil {
		log.Fatalf("failed saved session %v", err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	// gen authenticator and callback
	authenticator, err := auth.NewAuthenticator()
	if err != nil {
		log.Fatalf("failed gen authenticator", err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	authURL := authenticator.Config.AuthCodeURL(state)

	return c.Redirect(http.StatusTemporaryRedirect, authURL)
}