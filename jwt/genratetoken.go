package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	signingKey = []byte("secret") // Change this to your secret key
)

type CustomClaims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateToken() (string, error) {
	claims := CustomClaims{
		111111111,
		"bbaabbbuu",
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
			Issuer:    "your-issuer",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(signingKey)
}
