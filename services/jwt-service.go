package service

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTService interface {
	Generate() (string, error)
}

type jwtService struct {
	signingKey string //could be env variable in feature
}

func CreateJWTService(key string) JWTService {
	return &jwtService{signingKey: key}
}

// TODO: Think about client side validation for generated token
func (service *jwtService) Generate() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["user"] = "Jack Hunter"
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()

	tokenString, err := token.SignedString(service.signingKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
