package repository

import (
	"database/sql"

	"github.com/gasBlar/GoGoManager/models"
)

type UserRepository interface {
	FindById(id int) (models.ProfileManagerAuth, error)
	UpdatePartial(id int, user models.ProfileManagerUpdateRequest) (models.ProfileManagerAuth, error)
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
	err := r.db.QueryRow(query, id).Scan(&user.Id, &user.Email, &user.AuthId, &user.Name, &user.UserImageUri, &user.CompanyName, &user.CompanyImageUri)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) UpdatePartial(id int, user models.ProfileManagerUpdateRequest) (models.ProfileManagerAuth, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return models.ProfileManagerAuth{}, err
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		}
	}()

	if user.Name != "" || user.UserImageUri != "" || user.CompanyName != "" || user.CompanyImageUri != "" {
		query := "UPDATE profileManager SET"
		args := []interface{}{}

		if user.Name != "" {
			query += " name = ?,"
			args = append(args, user.Name)
		}
		if user.UserImageUri != "" {
			query += " userImage = ?,"
			args = append(args, user.UserImageUri)
		}
		if user.CompanyName != "" {
			query += " companyName = ?,"
			args = append(args, user.CompanyName)
		}
		if user.CompanyImageUri != "" {
			query += " companyImage = ?,"
			args = append(args, user.CompanyImageUri)
		}

		query = query[:len(query)-1] + " WHERE id = ?"
		args = append(args, id)

		_, err = tx.Exec(query, args...)
		if err != nil {
			return models.ProfileManagerAuth{}, err
		}
	}

	if user.Email != "" {
		emailQuery := "UPDATE auth SET email = ? WHERE id = (SELECT authId FROM profileManager WHERE id = ?)"
		_, err = tx.Exec(emailQuery, user.Email, id)
		if err != nil {
			return models.ProfileManagerAuth{}, err
		}
	}

	err = tx.Commit()
	if err != nil {
		return models.ProfileManagerAuth{}, err
	}

	updatedUser, err := r.FindById(id)
	if err != nil {
		return models.ProfileManagerAuth{}, err
	}

	return updatedUser, nil
}
