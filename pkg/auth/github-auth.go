package auth

import (
	"strings"

	"foundation/pkg/config"

	"github.com/go-pkgz/auth"
	"github.com/go-pkgz/auth/avatar"
	"github.com/go-pkgz/auth/token"
)

func NewGithubAuth(config config.Config) *auth.Service {
	options := auth.Opts{
		SecretReader: token.SecretFunc(func(id string) (string, error) { // secret key for JWT
			return config.JwtSecretKey, nil
		}),
		TokenDuration:  config.Auth.TokenDuration,
		CookieDuration: config.Auth.CookieDuration,
		Issuer:         config.WebServer.Name,
		URL:            config.WebServerURL(),
		AvatarStore:    avatar.NewLocalFS("/tmp"),
		Validator: token.ValidatorFunc(func(_ string, claims token.Claims) bool {
			// allow only dev_* names
			return claims.User != nil && strings.HasPrefix(claims.User.Name, "dev_")
		}),
	}

	// create auth service with providers
	service := auth.NewService(options)
	service.AddProvider("github", config.GithubClientID, config.GithubClientSecret) // add github provider
	return service
}
