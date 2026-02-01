package repository

import (
	"database/sql"

	"github.com/ahmadzakyarifin/echo-pos-admin/internal/domain"
	"github.com/jmoiron/sqlx"
)

type AuthRepository interface {
	FindUserByIdentifier(username, email string) (*domain.User, error)
}

type authRepository struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) AuthRepository {
	return &authRepository{db: db}
}

func (r *authRepository) FindUserByIdentifier(username, email string) (*domain.User, error) {
	var user domain.User
	query := "SELECT id, name, username, email, password, role FROM users WHERE username = ? OR email = ?"
	err := r.db.Get(&user, query, username, email)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &user, err
}
