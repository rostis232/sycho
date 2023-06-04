package service

import (
	"github.com/rostis232/psycho/internal/app/repository"
	"github.com/rostis232/psycho/internal/models"
)


type Authorization interface {
	CreateAndSaveUUID(login, password string) (string, error)
	GetUserByUUID(uuid string) (models.User, error)
	GetUsersOrganisation(orgID int) (models.Organisation, error)
	DeleteUUID(uuid string) error
}

type Beneficiary interface {
	GetAllBeneficiariesByUserID (userID int) ([]models.Beneficiary, error)
	GetBeneficiaryByID(benID int) (models.Beneficiary, error)
	GetActivitiesByBnfID(benID int) ([]models.Activity, error)
}

type Activity interface{}

type Organisation interface {
	GetAllOrganisations() ([]models.Organisation, error)
	AddOrganisation(title, code string) (int, error)
}

type Project interface {
	GetAllProjects() ([]models.Project, error)
}

type Service struct {
	Authorization
	Beneficiary
	Activity
	Organisation
	Project
}

func NewService (repo repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
		Beneficiary: NewBenService(repo.Beneficiary),
		Activity: NewActivityService(repo.Activity),
		Organisation: NewOrganisationService(repo.Organisation),
		Project: NewProjectService(repo.Project),
	}
}
