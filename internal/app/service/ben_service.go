package service

import (
	"github.com/rostis232/psycho/internal/app/repository"
	"github.com/rostis232/psycho/internal/models"
)

type BenService struct {
	repo repository.Beneficiary
}

func NewBenService (repo repository.Beneficiary) *BenService {
	return &BenService{
		repo: repo,
	}
}

func (b *BenService) GetAllBeneficiariesByUserID (userID int) ([]models.Beneficiary, error) {
	return b.repo.GetAllBeneficiariesByUserID(userID)
}

func (b *BenService) GetBeneficiaryByID(benID int) (models.Beneficiary, error) {
	return b.repo.GetBeneficiaryByID(benID)
}

func (b *BenService) GetActivitiesByBnfID(benID int) ([]models.Activity, error) {
	return b.repo.GetActivitiesByBnfID(benID)
}