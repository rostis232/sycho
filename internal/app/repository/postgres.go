package repository

import (
	"fmt"
	"log"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	sessionsTable = "sessions"
	usersTable = "users"
	projectsTable = "projects"
	organisationsTable = "organisations"
	beneficiariesTable = "beneficiaries"
	requestsTable = "reqs"
	activitiesTable = "activities"
	projectOrganizationTable = "prj_org"
	projectUserTable = "prj_usr"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	log.Printf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
	cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode)
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}