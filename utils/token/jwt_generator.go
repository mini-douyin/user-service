package token

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var jwtKey = []byte("YV9zZWNyZXRfa2V5") // TODO: put it in safe

func GenerateToken(userId uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
