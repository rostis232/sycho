package service

import (
	"github.com/rostis232/psycho/internal/app/repository"
	"github.com/rostis232/psycho/internal/models"
)

type OrganisationService struct {
	repo repository.Organisation
}

func NewOrganisationService (repo repository.Organisation) *OrganisationService {
	return &OrganisationService{
		repo: repo,
	}
}

func (o *OrganisationService) GetAllOrganisations() ([]models.Organisation, error) {
	orgs, err := o.repo.GetAllOrganisations()
	if err != nil {
		return orgs, err
	}
	for i := range orgs {
		orgs[i] = models.CheckOrgForNils(orgs[i])
	}
	return orgs, nil
}

func (o *OrganisationService) AddOrganisation(title, code string) (int, error) {
	org := models.Organisation{
		Title: &title,
		Code: &code,
	}
	return o.repo.AddOrganisation(org)
}