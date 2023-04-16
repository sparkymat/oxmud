package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sparkymat/oxmud/view"
)

func Home() echo.HandlerFunc {
	return func(c echo.Context) error {
		csrfToken := getCSRFToken(c)
		if csrfToken == "" {
			log.Print("error: csrf token not found")

			//nolint:wrapcheck
			return c.String(http.StatusInternalServerError, "server error")
		}

		pageHTML := view.Home()
		htmlString := view.Layout("oxmud", csrfToken, pageHTML)

		//nolint:wrapcheck
		return c.HTMLBlob(http.StatusOK, []byte(htmlString))
	}
}
