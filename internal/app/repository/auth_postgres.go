package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/rostis232/psycho/internal/models"
)

type AuthorizationPostgres struct {
	db *sqlx.DB
}

func NewAuthorizationPostgres (db *sqlx.DB) *AuthorizationPostgres {
	return &AuthorizationPostgres{
		db: db,
	}
}

func (a *AuthorizationPostgres) GetUser(login, password string) (models.User, error) {
	var user models.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE login=$1 AND pass=$2", usersTable)
	err := a.db.Get(&user, query, login, password)
	return user, err
}

func (a *AuthorizationPostgres) CreateAndSaveUUID(userID int) (string, error) {
	var uuid string
	query := fmt.Sprintf("INSERT INTO %s (user_id) values ($1) RETURNING id", sessionsTable)
	row := a.db.QueryRow(query, userID)
	if err := row.Scan(&uuid); err != nil {
		return "", err
	}
	return uuid, nil
}

func (a *AuthorizationPostgres) GetUserByUUID(uuid string) (models.User, error) {
	var user models.User
	query := fmt.Sprintf("SELECT u.user_id, u.login, u.pass, u.first_name, u.last_name, u.email," +
	" u.phone, u.org_id, u.role FROM %s u, %s s WHERE u.user_id=s.user_id AND s.id=$1", usersTable, sessionsTable)
	err := a.db.Get(&user, query, uuid)
	return user, err
}

func (a *AuthorizationPostgres) DeleteUUID(uuid string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", sessionsTable)
	sql, err := a.db.Exec(query, uuid)
	fmt.Println(sql)

	return err
}