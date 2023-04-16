package route

import (
	"context"

	"github.com/sparkymat/oxmud/dbx"
)

type ConfigService interface {
	JWTSecret() string
	SessionSecret() string
	DatabaseURL() string
	DisableRegistration() bool
	ReverseProxyAuthentication() bool
	ProxyAuthUsernameHeader() string
}

type DatabaseService interface {
	FetchUserByUsername(ctx context.Context, email string) (dbx.User, error)
	CreateUser(ctx context.Context, arg dbx.CreateUserParams) (dbx.User, error)
}
