package logout

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"net/url"
	"os"
)

func Handler(c echo.Context) error {
	domain := os.Getenv("AUTH0_DOMAIN")
	logoutUrl, err := url.Parse("https://" + domain)
	if err != nil {
		log.Fatalf("failed log out url parsed %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	logoutUrl.Path += "/v2/logout"
	parameters := url.Values{}

	returnUrl, err := url.Parse("http://localhost:8080")
	if err != nil {
		log.Fatalf("failed log out url parsed %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	parameters.Add("returnTo", returnUrl.String())
	parameters.Add("client_id", os.Getenv("AUTH0_CLIENT_ID"))
	logoutUrl.RawQuery = parameters.Encode()

	return c.Redirect(http.StatusTemporaryRedirect, logoutUrl.String())
}