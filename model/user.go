package model

import (
	"github.com/sally0226/oidc-go-example/types"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email          string         `gorm:"size:30;not null""`
	Name           string         `gorm:"size:30;not null""`
	Picture        string         `gorm:"size:255;not null""`
	Provider       types.Provider `gorm:"size:30;not null;uniqueIndex:ix_provider_provider_user_id"`
	ProviderUserID string         `gorm:"size:255;not null;uniqueIndex:ix_provider_provider_user_id"`
}
