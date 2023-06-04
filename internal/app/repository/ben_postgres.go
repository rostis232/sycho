package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/rostis232/psycho/internal/models"
)

type BeneficiaryPostgres struct {
	db *sqlx.DB
}

func NewBeneficiaryPostgres(db *sqlx.DB) *BeneficiaryPostgres {
	return &BeneficiaryPostgres{
		db: db,
	}
}

func (b *BeneficiaryPostgres) GetAllBeneficiariesByUserID(userID int) ([]models.Beneficiary, error) {
	bens := []models.Beneficiary{}
	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id=$1", beneficiariesTable)
	if err := b.db.Select(&bens, query, userID); err != nil {
		return nil, err
	}
	return bens, nil
}

func (b *BeneficiaryPostgres) GetBeneficiaryByID(benID int) (models.Beneficiary, error) {
	ben := models.Beneficiary{}
	query := fmt.Sprintf("SELECT * FROM %s WHERE bnf_id=$1", beneficiariesTable)
	err := b.db.Get(&ben, query, benID)
	return ben, err
}

func (b *BeneficiaryPostgres) GetActivitiesByBnfID(benID int) ([]models.Activity, error) {
	acts := []models.Activity{}
	query := fmt.Sprintf("SELECT * FROM %s WHERE bnf_id=$1", activitiesTable)
	if err := b.db.Select(&acts, query, benID); err != nil {
		return nil, err
	}
	return acts, nil
}