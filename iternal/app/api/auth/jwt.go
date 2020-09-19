package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecretKey = []byte("oblique")

// CreateJWT creates jwt token when user sign in
func CreateJWT(email string) (response string, err error) {
	expirationTime := time.Now().Add(10 * time.Minute)
	claims := &Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// VerifyToken will be used when need to verify token in endpoints
func VerifyToken(tokenString string) (email string, err error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecretKey, nil
	})

	if token != nil {
		return claims.Email, nil
	}
	return "", err
}
