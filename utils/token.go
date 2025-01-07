package utils

import (
	"time"

	"github.com/gasBlar/GoGoManager/config"
	"github.com/gasBlar/GoGoManager/models"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey []byte

type Claims struct {
	Email     string `json:"email"`
	ManagerId int    `json:"managerId"`
	jwt.RegisteredClaims
}

func init() {
	config.InitEnv()
	secretKey = []byte(config.GetEnv("SECRET_KEY"))
}

func CreateToken(userProfile models.ProfileManagerClaims) (string, error) {
	claims := Claims{
		Email:     userProfile.Email,
		ManagerId: userProfile.ManagerId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 2)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		}}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) (any, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)

	if !ok || !token.Valid {
		return nil, err
	}

	return claims, nil
}
