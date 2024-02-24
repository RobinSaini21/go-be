package jwt

import (
	"time"

	"begain.com/structutil"
	"github.com/dgrijalva/jwt-go"
)

var (
	signingKey = []byte("secret") // Change this to your secret key
)

type RequestBody struct {
	Username     string `bson:"username"`
	Email        string `bson:"email"`
	ReferralCode string `bson:"referral_code"`
	Password     string `bson:"password"`
}

type CustomClaims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateToken(user structutil.RequestBody) (string, error) {
	claims := CustomClaims{
		111111111,
		user.Email,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
			Issuer:    "your-issuer",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(signingKey)
}
