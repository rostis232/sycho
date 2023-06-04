package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/rostis232/psycho/internal/models"
)

type OrganisationPostgres struct {
	db *sqlx.DB
}

func NewOrganisationPostgres (db *sqlx.DB) *OrganisationPostgres {
	return &OrganisationPostgres{
		db: db,
	}
}

func (o *OrganisationPostgres) GetAllOrganisations() ([]models.Organisation, error) {
	organisations := []models.Organisation{}
	query := fmt.Sprintf("SELECT * FROM %s", organisationsTable)
	if err := o.db.Select(&organisations, query); err != nil {
		return nil, err
	}
	return organisations, nil
}

func (o *OrganisationPostgres) AddOrganisation(org models.Organisation) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (title, code) values ($1, $2) RETURNING org_id", organisationsTable)
	row := o.db.QueryRow(query, org.Title, org.Code)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}