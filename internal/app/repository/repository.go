package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/rostis232/psycho/internal/models"
)

type Authorization interface {
	GetUser(login, password string) (models.User, error)
	CreateAndSaveUUID(userID int) (string, error)
	GetUserByUUID(uuid string) (models.User, error)
	DeleteUUID(uuid string) error
}

type Beneficiary interface {
	GetAllBeneficiariesByUserID(userID int) ([]models.Beneficiary, error)
	GetBeneficiaryByID(benID int) (models.Beneficiary, error)
	GetActivitiesByBnfID(benID int) ([]models.Activity, error)
}

type Activity interface {

}

type Repository struct {
	Authorization
	Beneficiary
	Activity
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthorizationPostgres(db),
		Beneficiary:      NewBeneficiaryPostgres(db),
		Activity:      NewActivityPostgres(db),
	}
}