package repository

import (
	"database/sql"

	"github.com/gasBlar/GoGoManager/models"
)

type UserRepository interface {
	FindById(id int) (models.ProfileManagerAuth, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	if db == nil {
		panic("database connection cannot be nil")
	}
	return &userRepository{db: db}
}

func (r *userRepository) FindById(id int) (models.ProfileManagerAuth, error) {
	var user models.ProfileManagerAuth

	query := "SELECT pm.id, a.email, pm.authId, pm.name, pm.userImage, pm.companyName, pm.companyImage FROM profileManager pm LEFT JOIN auth a on pm.authId = a.id WHERE pm.id = ?"
	err := r.db.QueryRow(query, id).Scan(&user.Id, &user.Email, &user.AuthId, &user.Name, &user.UserImage, &user.CompanyName, &user.CompanyImage)

	if err != nil {
		return user, err
	}

	return user, nil
}
