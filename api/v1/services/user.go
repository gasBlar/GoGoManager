package services

import (
	"fmt"

	"github.com/gasBlar/GoGoManager/db"
	"github.com/gasBlar/GoGoManager/models"
	"github.com/gasBlar/GoGoManager/repository"
)

func GetUserProfile(id int) (models.ProfileManagerResponse, error) {
	database := db.DB

	userRepo := repository.NewUserRepository(database)

	user, err := userRepo.FindById(id)
	if err != nil {
		return models.ProfileManagerResponse{}, err
	}

	result := user.ToResponse()

	return result, nil
}

func UpdateUserProfile(id int, user models.ProfileManagerUpdateRequest) (models.ProfileManagerResponse, error) {
	database := db.DB

	authRepo := repository.NewAuthRepository(database)
	isExist := authRepo.CheckEmailExistExceptManagerId(user.Email, id)
	if isExist {
		return models.ProfileManagerResponse{}, fmt.Errorf("email already exists")
	}

	userRepo := repository.NewUserRepository(database)
	userUpdated, err := userRepo.UpdatePartial(id, user)
	if err != nil {
		return models.ProfileManagerResponse{}, err
	}

	result := userUpdated.ToResponse()

	return result, nil
}
