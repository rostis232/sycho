package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/rostis232/psycho/internal/models"
)

type ProjectPostgres struct {
	db *sqlx.DB
}

func NewProjectPostgres (db *sqlx.DB) *ProjectPostgres {
	return &ProjectPostgres{
		db: db,
	}
}

func (p *ProjectPostgres) GetAllProjects() ([]models.Project, error) {
	projects := []models.Project{}
	query := fmt.Sprintf("SELECT * FROM %s", projectsTable)
	if err := p.db.Select(&projects, query); err != nil {
		return nil, err
	}
	return projects, nil
}