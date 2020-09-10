package service

import (
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"time"
)

type JWTService interface {
	GenerateToken() string
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtServices struct {
	secretKey string
	issure    string
}

func JWTAuthService() JWTService {
	return &jwtServices{
		secretKey: "secret",
		issure:    "Daniel",
	}
}

func (service *jwtServices) GenerateToken() string {
	claims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
		Issuer:    service.issure,
		IssuedAt:  time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(service.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

func (service *jwtServices) ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("Invalid token", token.Header["alg"])

		}
		return []byte(service.secretKey), nil
	})

}
