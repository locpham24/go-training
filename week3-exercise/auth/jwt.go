package auth

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JwtClaims struct {
	Name string `json:"name"`
	jwt.StandardClaims
}

func CreateJwtToken(name string, id string) (string, error) {
	jwtClaims := JwtClaims{
		name,
		jwt.StandardClaims{
			Id:        id,
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}
	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS512, jwtClaims)
	token, err := rawToken.SignedString([]byte(os.Getenv("JWT_KEY")))
	if err != nil {
		return "", err
	}
	return token, nil
}
