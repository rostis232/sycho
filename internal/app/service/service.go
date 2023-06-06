package service

import (
	"github.com/rostis232/psycho/internal/app/repository"
	"github.com/rostis232/psycho/internal/models"
)


type Authorization interface {
	CreateAndSaveUUID(login, password string) (string, error)
	GetUserByUUID(uuid string) (models.User, error)
	DeleteUUID(uuid string) error
	NewUserRegistration (login, pass, fName, lName, email, phone string) error
}

type Beneficiary interface {
	GetAllBeneficiariesByUserID (userID int) ([]models.Beneficiary, error)
	GetBeneficiaryByID(benID int) (models.Beneficiary, error)
	GetActivitiesByBnfID(benID int) ([]models.Activity, error)
}

type Activity interface{}


type Service struct {
	Authorization
	Beneficiary
	Activity
}

func NewService (repo repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
		Beneficiary: NewBenService(repo.Beneficiary),
		Activity: NewActivityService(repo.Activity),
	}
}
