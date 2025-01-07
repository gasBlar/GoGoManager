package repository

import (
	"database/sql"

	"github.com/gasBlar/GoGoManager/models"
)

type AuthRepository interface {
	FindByEmail(email string) (models.Auth, error)
}

type authRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) AuthRepository {
	if db == nil {
		panic("database connection cannot be nil")
	}
	return &authRepository{db: db}
}

func (r *authRepository) FindByEmail(email string) (models.Auth, error) {
	var auth models.Auth

	query := "SELECT Id, email, password FROM auth WHERE email = ?"
	err := r.db.QueryRow(query, email).Scan(&auth.Id, &auth.Email, &auth.Password)

	if err != nil {
		return auth, err
	}

	return auth, nil
}
