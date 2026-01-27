package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/zmonsoonz/go-test-api/internal/content/domain"
	"github.com/zmonsoonz/go-test-api/internal/platform/db/postgres"
)
type AuthRep struct {
	db *sqlx.DB
}

func NewAuthRep(db *sqlx.DB) *AuthRep {
	return &AuthRep{db: db}
}

func (r *AuthRep) CreateUser(user domain.User) (int, error)  {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (username, password_hash, email, created_at) VALUES ($1, $2, $3, $4) RETURNING user_id", postgres.UsersTable)

	row := r.db.QueryRow(query, user.Username, user.Password, user.Email, user.CreatedAt)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}