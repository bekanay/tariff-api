package repository

import (
	"context"
	"tariff-calculator/tariff-api/pkg/config"
	"tariff-calculator/tariff-api/pkg/model"

	"github.com/Nerzal/gocloak/v11"
	"github.com/pkg/errors"
)

type AuthRepository struct {
	client *gocloak.Client
	cfg    *config.Config
}

func NewAuthRepository(cfg *config.Config) *AuthRepository {
	client := gocloak.NewClient(cfg.KeycloakURL)
	return &AuthRepository{
		client: client,
		cfg:    cfg,
	}
}

func (repo *AuthRepository) Login(username, password string) (*model.TokenResponse, error) {
	token, err := repo.client.Login(context.Background(), repo.cfg.ClientID, repo.cfg.ClientSecret, repo.cfg.Realm, username, password)
	if err != nil {
		return nil, errors.Wrap(err, "failed to login")
	}
	return &model.TokenResponse{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		IdToken:      token.IdToken,
		TokenType:    token.TokenType,
		ExpiresIn:    token.ExpiresIn,
	}, nil
}

func (repo *AuthRepository) RefreshToken(refreshToken string) (*model.TokenResponse, error) {
	token, err := repo.client.RefreshToken(context.Background(), repo.cfg.ClientID, refreshToken, repo.cfg.Realm)
	if err != nil {
		return nil, errors.Wrap(err, "failed to refresh token")
	}
	return &model.TokenResponse{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		IdToken:      token.IdToken,
		TokenType:    token.TokenType,
		ExpiresIn:    token.ExpiresIn,
	}, nil
}

func (repo *AuthRepository) Logout(refreshToken string) error {
	err := repo.client.Logout(context.Background(), repo.cfg.ClientID, refreshToken, repo.cfg.Realm)
	if err != nil {
		return errors.Wrap(err, "failed to logout")
	}
	return nil
}
