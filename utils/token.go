package utils

import (
	"time"

	"github.com/gasBlar/GoGoManager/config"
	"github.com/gasBlar/GoGoManager/models"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey []byte

func init() {
	config.InitEnv()
	secretKey = []byte(config.GetEnv("SECRET_KEY"))
}

func CreateToken(userProfile models.ProfileManagerClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"email":     userProfile.Email,
			"managerId": userProfile.ManagerId,
			"exp":       time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}
