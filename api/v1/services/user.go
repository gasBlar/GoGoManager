package services

import (
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
