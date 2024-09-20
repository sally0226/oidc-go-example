package service

import "github.com/sally0226/oidc-go-example/types"

type IOAuthService interface {
	AuthURL() string
	ExchangeToken(code string) (string, error)
	ParseUser(idToken string) (*OAuthUser, error)
}

type OAuthUser struct {
	ID      string
	Email   string
	Name    string
	Picture string
}

func NewOAuthServices() map[types.Provider]IOAuthService {
	return map[types.Provider]IOAuthService{
		types.ProviderGoogle: newGoogleOAuthService(),
	}
}
