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

func (r *AuthRep) Create(user domain.User) (int, error)  {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()
	var id int

	createUserQuery := fmt.Sprintf("INSERT INTO %s (username, password_hash, email, created_at) VALUES ($1, $2, $3, $4) RETURNING id", postgres.UsersTable)
	row := tx.QueryRow(createUserQuery, user.Username, user.Password, user.Email, user.CreatedAt)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	createLibraryQuery := fmt.Sprintf("INSERT INTO %s (user_id, created_at) VALUES ($1, $2)", postgres.LibrariesTable)
	_, err = tx.Exec(createLibraryQuery, id, user.CreatedAt)

	if err != nil {
		tx.Rollback()
		return 0, err
	}
	return id, tx.Commit()
}


func (r *AuthRep) GetByUsername(username string) (domain.User, error)  {
	var user domain.User
	query := fmt.Sprintf("SELECT id, password_hash FROM %s WHERE username = $1", postgres.UsersTable)

	err := r.db.Get(&user, query, username)

	return user, err
}