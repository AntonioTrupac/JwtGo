package service

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type JwtCustomClaim struct {
	ID int `json:"id"`
	jwt.StandardClaims
}

var jwtSecret = []byte(getJwtSecret())

func getJwtSecret() string {
	secret := os.Getenv("JWT_SECRET")

	if secret == "" {
		return "aSecret"
	}

	return secret
}

func GenerateJwt(ctx context.Context, userID int) (string, error) {
	// create a new token with custom claims that is passed as a second param
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, &JwtCustomClaim{
		ID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	})

	// after the token is created, sign in with JWT_SECRET to make it secure
	token, err := tokenClaims.SignedString(jwtSecret)

	if err != nil {
		return "", err
	}

	return token, nil
}

func ValidateJwt(ctx context.Context, token string) (*jwt.Token, error) {
	// validate and return the token as *jwt.Token
	// first we check the token method
	// check if the signing methjod is HMAC since we use a secret HS256 to sign the token
	return jwt.ParseWithClaims(token, &JwtCustomClaim{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there has been a problem with the signing method")
		}

		return jwtSecret, nil
	})
}
