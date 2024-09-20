package types

import "errors"

type Provider string

const (
	ProviderGoogle Provider = "google"
	ProviderKakao           = "kakao"
)

func ValidateProvider(p string) (Provider, error) {
	provider := Provider(p)
	switch provider {
	case ProviderGoogle, ProviderKakao:
		return provider, nil
	default:
		return "", errors.New("not found provider")
	}
}
