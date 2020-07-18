package service

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTService interface {
	Generate() (string, error)
	Validate(tokenString string) (bool, error)
}

type jwtService struct {
	signingKey string
}

// Creates a JWT service which can generate a token and validate a token based on its key
func CreateJWTService(key string) JWTService {
	return &jwtService{signingKey: key}
}

func (service *jwtService) Generate() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["user"] = "Jack Hunter"
	claims["exp"] = time.Now().Add(time.Hour).Unix()

	tokenString, err := token.SignedString([]byte(service.signingKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (service *jwtService) Validate(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Unexpected signing method")
		}

		return []byte(service.signingKey), nil
	})

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		if claims["authorized"] == true && claims["user"] == "Jack Hunter" {
			return true, nil
		} else {
			return false, errors.New("Token is invalid")
		}
	}
	return false, err
}
