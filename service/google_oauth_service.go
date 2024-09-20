package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/sally0226/oidc-go-example/types"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"os"
)

type googleOAuthService struct {
	provider types.Provider
	config   *oauth2.Config
}

func (s googleOAuthService) ExchangeToken(code string) (string, error) {
	token, err := s.config.Exchange(context.Background(), code)
	if err != nil {
		return "", fmt.Errorf("googleOAuthService.ExchangeToken : %w", err)
	}

	idToken, ok := token.Extra("id_token").(string)
	if !ok {
		return "", errors.New("googleOAuthService.ExchangeToken : no id token")
	}

	return idToken, nil
}

func (s googleOAuthService) ParseUser(token string) (*OAuthUser, error) {
	provider, err := oidc.NewProvider(context.Background(), "https://accounts.google.com")
	if err != nil {
		return nil, fmt.Errorf("googleOAuthService.IsMember : %w", err)
	}
	verifier := provider.Verifier(&oidc.Config{ClientID: s.config.ClientID})

	// Parse and verify ID Token payload.
	idToken, err := verifier.Verify(context.Background(), token)
	if err != nil {
		return nil, fmt.Errorf("googleOAuthService.IsMember : %w", err)
	}

	// Extract custom claims
	var claims struct {
		Iss           string `json:"iss"`
		Azp           string `json:"azp"`
		Aud           string `json:"aud"` // client id
		Sub           string `json:"sub"` // provider user id
		Email         string `json:"email"`
		EmailVerified bool   `json:"email_verified"`
		AtHash        string `json:"at_hash"`
		Name          string `json:"name"`
		Picture       string `json:"picture"`
		GivenName     string `json:"given_name"`
		FamilyName    string `json:"family_name"`
		Iat           int    `json:"iat"`
		Exp           int    `json:"exp"`
	}
	if err = idToken.Claims(&claims); err != nil {
		return nil, err
	}

	return &OAuthUser{
		ID:      claims.Sub,
		Email:   claims.Email,
		Name:    claims.Name,
		Picture: claims.Picture,
	}, nil
}

func (s googleOAuthService) AuthURL() string {
	// TODD: state 적용
	return s.config.AuthCodeURL("")

}

func newGoogleOAuthService() *googleOAuthService {
	return &googleOAuthService{
		provider: types.ProviderGoogle,
		config: &oauth2.Config{
			ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
			ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
			Endpoint:     google.Endpoint,
			RedirectURL:  os.Getenv("OAUTH_CALLBACK_URL"),
			//Scopes:       []string{"openid"}, // 이메일, 프로필 등 부가정보를 받을 수 없음
			Scopes: []string{"openid", "profile", "email"},
		},
	}
}
