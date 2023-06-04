package service

import (
	"github.com/rostis232/psycho/internal/app/repository"
	"github.com/rostis232/psycho/internal/models"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService (repo repository.Authorization) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (a *AuthService) CreateAndSaveUUID(login, password string) (string, error) {
	user, err := a.repo.GetUser(login, password)
	if err != nil {
		return "", err
	}

	uuid, err := a.repo.CreateAndSaveUUID(*user.UserId)
	if err != nil {
		return "", err
	}

	return uuid, nil
}

func (a *AuthService) GetUserByUUID(uuid string) (models.User, error) {
	user, err := a.repo.GetUserByUUID(uuid)
	if err != nil {
		return user, err
	}
	return models.CheckUserForNils(user), nil
}

func (a *AuthService) GetUsersOrganisation(orgID int) (models.Organisation, error) {
	org, err := a.repo.GetUsersOrganisation(orgID)
	if err != nil {
		return org, err
	}
	return models.CheckOrgForNils(org), nil
}

func (a *AuthService) DeleteUUID(uuid string) error {
	return a.repo.DeleteUUID(uuid)
}
