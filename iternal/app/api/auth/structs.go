package auth

import "github.com/dgrijalva/jwt-go"

type RegistationParams struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Claims struct {
	Email string
	jwt.StandardClaims
}
