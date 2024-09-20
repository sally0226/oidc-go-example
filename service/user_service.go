package service

import (
	"github.com/sally0226/oidc-go-example/model"
	"github.com/sally0226/oidc-go-example/repository"
	"github.com/sally0226/oidc-go-example/types"
)

type IUserService interface {
	CreateUser(u *model.User) (*model.User, error)
	GetUserByProvider(p types.Provider, puid string) (*model.User, error)
}

type userService struct {
	ur repository.IUserRepository
}

func (s *userService) GetUserByProvider(p types.Provider, puid string) (*model.User, error) {
	return s.ur.GetUserByProvider(p, puid)
}

func (s *userService) CreateUser(u *model.User) (*model.User, error) {
	if err := s.ur.CreateUser(u); err != nil {
		return nil, err
	}
	return s.ur.GetUserByProvider(u.Provider, u.ProviderUserID)
}

func NewUserService(ur repository.IUserRepository) IUserService {
	return &userService{
		ur: ur,
	}
}
