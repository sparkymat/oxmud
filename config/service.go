package config

import (
	"fmt"
	"net/url"

	"github.com/caarlos0/env/v7"
)

func New() (*Service, error) {
	var envValues envValues

	err := env.Parse(&envValues)
	if err != nil {
		return nil, fmt.Errorf("failed to load config from env. err: %w", err)
	}

	return &Service{
		envValues: envValues,
	}, nil
}

type Service struct {
	envValues envValues
}

type envValues struct {
	JWTSecret                  string `env:"JWT_SECRET,required"`
	SessionSecret              string `env:"SESSION_SECRET,required"`
	DatabaseName               string `env:"DATABASE_NAME,required"`
	DatabaseHostname           string `env:"DATABASE_HOSTNAME,required"`
	DatabasePort               string `env:"DATABASE_PORT,required"`
	DatabaseUsername           string `env:"DATABASE_USERNAME"`
	DatabasePassword           string `env:"DATABASE_PASSWORD"`
	DatabaseSSLMode            bool   `env:"DATABASE_SSL_MODE" envDefault:"true"`
	DisableRegistration        bool   `env:"DISABLE_REGISTRATION" envDefault:"false"`
	ReverseProxyAuthentication bool   `env:"REVERSE_PROXY_AUTHENTICATION" envDefault:"false"`
	ProxyAuthUsernameHeader    string `env:"PROXY_AUTH_USERNAME_HEADER" envDefault:"Remote-User"`
}

func (s *Service) JWTSecret() string {
	return s.envValues.JWTSecret
}

func (s *Service) SessionSecret() string {
	return s.envValues.SessionSecret
}

func (s *Service) DatabaseURL() string {
	connString := "postgres://"

	if s.envValues.DatabaseUsername != "" {
		connString = fmt.Sprintf("%s%s", connString, s.envValues.DatabaseUsername)

		if s.envValues.DatabasePassword != "" {
			encodedPassword := url.QueryEscape(s.envValues.DatabasePassword)
			connString = fmt.Sprintf("%s:%s", connString, encodedPassword)
		}

		connString = fmt.Sprintf("%s@", connString)
	}

	sslMode := "disable"
	if s.envValues.DatabaseSSLMode {
		sslMode = "require"
	}

	connString = fmt.Sprintf(
		"%s%s:%s/%s?sslmode=%s",
		connString,
		s.envValues.DatabaseHostname,
		s.envValues.DatabasePort,
		s.envValues.DatabaseName,
		sslMode,
	)

	return connString
}

func (s *Service) DisableRegistration() bool {
	return s.envValues.DisableRegistration
}

func (s *Service) ReverseProxyAuthentication() bool {
	return s.envValues.ReverseProxyAuthentication
}

func (s *Service) ProxyAuthUsernameHeader() string {
	return s.envValues.ProxyAuthUsernameHeader
}
