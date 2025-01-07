package services

import (
	"database/sql"
	"fmt"

	"github.com/gasBlar/GoGoManager/models"
	"github.com/gasBlar/GoGoManager/utils"
)

var db *sql.DB

func Login(auth models.AuthLoginRequest) (models.AuthLoginResponse, error) {

	// Check if the user exists
	// var email string
	// err := db.QueryRow("SELECT email FROM auth WHERE email = ? AND password = ?", auth.Email, auth.Password).Scan(&email)

	var result models.AuthLoginResponse
	// if err != nil {
	// 	return result, err
	// } else {
	token, err := utils.CreateToken(models.ProfileManagerClaims{ManagerId: 1, Email: auth.Email})
	if err != nil {
		return models.AuthLoginResponse{}, err
	}

	result = models.AuthLoginResponse{Email: auth.Email, Token: token}
	return result, nil
	// }
}

func Register(auth models.AuthLoginRequest) (models.AuthLoginResponse, error) {
	hashedPassword, err := utils.HashPassword(auth.Password)
	if err != nil {
		return models.AuthLoginResponse{}, err
	}

	var existingEmail string
	err = db.QueryRow("SELECT email FROM auth WHERE email = ?", auth.Email).Scan(&existingEmail)
	if err != sql.ErrNoRows {
		return models.AuthLoginResponse{}, fmt.Errorf("email already exists")
	}

	_, err = db.Exec("INSERT INTO auth (email, password) VALUES (?, ?)", auth.Email, hashedPassword)
	if err != nil {
		return models.AuthLoginResponse{}, err
	}

	token, err := utils.CreateToken(models.ProfileManagerClaims{Email: auth.Email})

	if err != nil {
		return models.AuthLoginResponse{}, err
	}

	return models.AuthLoginResponse{Email: auth.Email, Token: token}, nil
}
