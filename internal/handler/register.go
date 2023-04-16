package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sparkymat/oxmud/dbx"
	"github.com/sparkymat/oxmud/view"
	"golang.org/x/crypto/bcrypt"
)

const (
	minPasswordLength = 8
	defaultBcryptCost = 10
)

func Register(_ ConfigService, _ DatabaseService) echo.HandlerFunc {
	return func(c echo.Context) error {
		return renderRegistrationPage(c, "", "")
	}
}

func DoRegister(_ ConfigService, db DatabaseService) echo.HandlerFunc {
	return func(c echo.Context) error {
		username := c.FormValue("username")
		password := c.FormValue("password")
		passwordConfirmation := c.FormValue("password_confirmation")

		if password != passwordConfirmation {
			log.Printf("passwords don't match")

			return renderRegistrationPage(c, username, "Passwords don't match")
		}

		if len(password) < minPasswordLength {
			log.Printf("password too short")

			return renderRegistrationPage(c, username, "Password too short")
		}

		_, err := db.FetchUserByUsername(c.Request().Context(), username)
		if err == nil {
			log.Printf("failed to load user. err: %v", err)

			return renderRegistrationPage(c, username, "User already registered")
		}

		encryptedPasswordBytes, err := bcrypt.GenerateFromPassword([]byte(password), defaultBcryptCost)
		if err != nil {
			log.Printf("password does not match")

			return renderRegistrationPage(c, username, "Registration failed")
		}

		_, err = db.CreateUser(c.Request().Context(), dbx.CreateUserParams{
			Username:          username,
			EncryptedPassword: string(encryptedPasswordBytes),
		})
		if err != nil {
			log.Printf("failed to create new user. err: %v", err)

			return renderRegistrationPage(c, username, "Registration failed")
		}

		//nolint:wrapcheck
		return c.Redirect(http.StatusSeeOther, "/login")
	}
}

func renderRegistrationPage(c echo.Context, username string, errorMessage string) error {
	csrfToken := getCSRFToken(c)
	if csrfToken == "" {
		log.Print("error: csrf token not found")

		//nolint:wrapcheck
		return c.String(http.StatusInternalServerError, "server error")
	}

	pageHTML := view.Register(csrfToken, username, errorMessage)
	htmlString := view.Layout("oxmud | register", csrfToken, pageHTML)

	//nolint:wrapcheck
	return c.HTMLBlob(http.StatusOK, []byte(htmlString))
}
