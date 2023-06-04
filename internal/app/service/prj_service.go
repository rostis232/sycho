package service

import (
	"github.com/rostis232/psycho/internal/app/repository"
	"github.com/rostis232/psycho/internal/models"
)

type ProjectService struct {
	repo repository.Project
}

func NewProjectService (repo repository.Project) *ProjectService {
	return &ProjectService{
		repo: repo,
	}
}

func (p *ProjectService) GetAllProjects() ([]models.Project, error) {
	prjs, err := p.repo.GetAllProjects()
	if err != nil {
		return prjs, err
	}
	for i := range prjs {
		prjs[i] = models.CheckProjectForNil(prjs[i])
	}
	return prjs, nil
}