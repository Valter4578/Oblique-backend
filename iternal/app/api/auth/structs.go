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

type response struct {
	Token   string `json:"token,omitempty"`
	Email   string `json:"email,omitempty"`
	Message string `json:"message,omitempty"`
}
