package webtoken

import (
	"time"
	"github.com/dgrijalva/jwt-go"
)

const secret string = "SUPER SECRET"

type WebToken struct {
	Claims jwt.MapClaims
	Valid bool
}

func New(subject string) (string, error) {
	// Build a signed JSON Web Token
	claims := jwt.StandardClaims{
		Subject: subject,
		IssuedAt: time.Now().Unix(),
		ExpiresAt: time.Now().Add(time.Second * 3600 * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func GetSubject(tokenString string) (sub string, err error) {
	token, err := jwt.Parse(tokenString, getSecret)
	if err != nil {
		return // Don't want to proceed if token parsing failed
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		sub, ok = claims["sub"].(string)
	}

	return 
}

func Parse(tokenString string) (token WebToken, err error) {
	parsed, err := jwt.Parse(tokenString, getSecret)

	token.Valid = parsed.Valid

	if claims, ok := parsed.Claims.(jwt.MapClaims); ok {
		token.Claims = claims
	}

	return
}

func getSecret(token *jwt.Token) (interface{}, error) {
	return []byte(secret), nil
}