package adapters

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jailtonjunior94/bookings/api/infrastructure/environments"
)

type IJwtAdapter interface {
	GenerateTokenJWT(id, email string) (r string, err error)
	ExtractClaims(tokenString string) (id *string, err error)
}

type JwtAdapter struct {
}

func NewJwtAdapter() IJwtAdapter {
	return &JwtAdapter{}
}

func (j *JwtAdapter) GenerateTokenJWT(id, email string) (r string, err error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["sub"] = id
	claims["email"] = email
	claims["exp"] = time.Now().Local().Add(time.Hour * 24 * time.Duration(1))

	t, err := token.SignedString([]byte(environments.JwtSecret))
	if err != nil {
		return "", err
	}
	return t, err
}

func (j *JwtAdapter) ExtractClaims(tokenString string) (id *string, err error) {
	tokenString = strings.Split(tokenString, " ")[1]
	hmacSecret := []byte(environments.JwtSecret)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return hmacSecret, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("Token inválido")
	}

	sub := fmt.Sprintf("%v", claims["sub"])
	return &sub, nil
}
