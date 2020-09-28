package auth

import (
	"errors"
	"io"
	"log"
	"net/http"
	"oblique/iternal/app/logger"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword generates hash for password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(bytes), err
}

// CheckPasswordHash compare hash with password
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		log.Println(err)
	}
	return err == nil
}

func WriteError(writer http.ResponseWriter, message string, status int) {
	err := errors.New(message)
	msg := logger.JSONError(err)
	io.WriteString(writer, msg)
	writer.WriteHeader(status)
}
