package repository

import (
	"database/sql"

	"github.com/gasBlar/GoGoManager/models"
)

type AuthRepository interface {
	FindByEmail(email string) (models.AuthLogin, error)
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

func (r *authRepository) FindByEmail(email string) (models.AuthLogin, error) {
	var auth models.AuthLogin

	query := "SELECT a.Id, a.email, a.password, pm.id as managerId  FROM auth a LEFT JOIN profileManager pm on a.id = pm.authId WHERE email = ?"
	err := r.db.QueryRow(query, email).Scan(&auth.Id, &auth.Email, &auth.Password, &auth.ProfileManagerId)

	if err != nil {
		return auth, err
	}

	return auth, nil
}
