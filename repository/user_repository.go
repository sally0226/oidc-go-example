package repository

import (
	"github.com/sally0226/oidc-go-example/model"
	"github.com/sally0226/oidc-go-example/types"
	"gorm.io/gorm"
)

type IUserRepository interface {
	GetUserByProvider(p types.Provider, providerUserID string) (*model.User, error)
	CreateUser(u *model.User) error
	GetUser(id string) (*model.User, error)
}
type userRepository struct {
	db *gorm.DB
}

func (r *userRepository) GetUserByProvider(p types.Provider, providerUserID string) (*model.User, error) {
	var u model.User
	if err := r.db.Where("provider = ?", p).
		Where("provider_user_id = ?", providerUserID).
		Find(&u).Error; err != nil {
		return nil, err
	}

	if u.ID == 0 {
		return nil, nil
	}

	return &u, nil
}

func (r *userRepository) CreateUser(u *model.User) error {
	if err := r.db.Create(u).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) GetUser(id string) (*model.User, error) {
	var u model.User
	if err := r.db.Find(&u, id).Error; err != nil {
		return nil, err
	}

	if u.ID == 0 {
		return nil, nil
	}

	return &u, nil
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}
