package services

import (
	"database/sql"
	"fmt"

	"github.com/gasBlar/GoGoManager/db"
	"github.com/gasBlar/GoGoManager/models"
	"github.com/gasBlar/GoGoManager/repository"
	"github.com/gasBlar/GoGoManager/utils"
)

func Login(auth models.AuthLoginRequest) (models.AuthLoginResponse, error) {
	database := db.DB

	authRepo := repository.NewAuthRepository(database)

	authData, err := authRepo.FindByEmail(auth.Email)
	if err != nil {
		return models.AuthLoginResponse{}, err
	}
	if authData.Email == "" {
		return models.AuthLoginResponse{}, fmt.Errorf("email not found")
	}

	var result models.AuthLoginResponse

	if err := utils.VerifyPassword(authData.Password, auth.Password); err != nil {
		return result, fmt.Errorf("invalid password")
	}

	token, err := utils.CreateToken(models.ProfileManagerClaims{Email: authData.Email})
	if err != nil {
		return models.AuthLoginResponse{}, err
	}

	result = models.AuthLoginResponse{Email: authData.Email, Token: token}
	return result, nil

}

func Register(auth models.AuthLoginRequest) (models.AuthLoginResponse, error) {
	database := db.DB
	tx, err := database.Begin()
	if err != nil {
		return models.AuthLoginResponse{}, err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()
	hashedPassword, err := utils.HashPassword(auth.Password)
	if err != nil {
		return models.AuthLoginResponse{}, err
	}

	var existingEmail string
	err = tx.QueryRow("SELECT email FROM auth WHERE email = ?", auth.Email).Scan(&existingEmail)
	if err != sql.ErrNoRows {
		return models.AuthLoginResponse{}, fmt.Errorf("email already exists")
	}

	result, err := tx.Exec("INSERT INTO auth (email, password) VALUES (?, ?)", auth.Email, hashedPassword)
	if err != nil {
		return models.AuthLoginResponse{}, err
	}
	if err != nil {
		return models.AuthLoginResponse{}, err
	}
	authID, err := result.LastInsertId()
	resProfile, err := tx.Exec("INSERT INTO profileManager (authId) VALUES (?)",
		authID)
	if err != nil {
		return models.AuthLoginResponse{}, err
	}
	profileId, err := resProfile.LastInsertId()

	if err = tx.Commit(); err != nil {
		return models.AuthLoginResponse{}, err
	}
	token, err := utils.CreateToken(models.ProfileManagerClaims{Id: int(profileId), AuthId: int(authID), Email: auth.Email})

	if err != nil {
		return models.AuthLoginResponse{}, err
	}

	return models.AuthLoginResponse{Email: auth.Email, Token: token}, nil
}
