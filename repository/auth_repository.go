package repository

import (
	"database/sql"

	"github.com/gasBlar/GoGoManager/models"
)

type AuthRepository interface {
	FindByEmail(email string) (models.AuthLogin, error)
	CheckEmailExist(email string) bool
	CheckEmailExistExceptManagerId(email string, id int) bool
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

func (r *authRepository) CheckEmailExist(email string) bool {
	var existingEmail string
	res := r.db.QueryRow("SELECT email FROM auth WHERE email = ?", email).Scan(&existingEmail)
	return res != sql.ErrNoRows
}

func (r *authRepository) CheckEmailExistExceptManagerId(email string, id int) bool {
	var existingEmail string
	res := r.db.QueryRow("SELECT email FROM auth WHERE email = ? AND id != ?", email, id).Scan(&existingEmail)
	return res != sql.ErrNoRows
}
