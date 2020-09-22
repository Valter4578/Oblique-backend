package auth

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"oblique/iternal/app/db"
	"oblique/iternal/app/model"
	"strings"

	"oblique/iternal/app/logger"
)

// error messages
const (
	missingEmail    = "Email is empty"
	missingPassword = "Password is empty"

	cantGetDataFromDb = "Couldn't get data from database"
	cantCreateUserDB  = "Couldn't create user in database"

	cantCreateJWT   = "Couldn't create JWT token"
	cantVerifyToken = "Couldn't verify token"
)

func SignIn(w http.ResponseWriter, r *http.Request) {
	log.Println("Auth: SignIn")

	var params LoginParams

	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		msg := logger.JSONError(err)
		io.WriteString(w, msg)
		return
	}

	if params.Email == "" {
		err = errors.New(missingEmail)
		msg := logger.JSONError(err)
		io.WriteString(w, msg)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if params.Password == "" {
		err = errors.New(missingPassword)
		msg := logger.JSONError(err)
		io.WriteString(w, msg)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := db.GetUser(params.Email, params.Password)
	if err != nil {
		err = errors.New(cantGetDataFromDb)
		msg := logger.JSONError(err)
		io.WriteString(w, msg)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tokenString, err := CreateJWT(params.Email)
	if err != nil {
		log.Println(err)
		err = errors.New(cantCreateJWT)
		msg := logger.JSONError(err)
		io.WriteString(w, msg)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	rspns := &response{
		Token: tokenString,
		Email: user.Email,
	}

	log.Println(rspns.Token)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rspns)
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	log.Println("Auth: SignUp")
	var params RegistationParams

	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		msg := logger.JSONError(err)
		io.WriteString(w, msg)
		return
	}

	if params.Email == "" {
		err = errors.New(missingEmail)
		msg := logger.JSONError(err)
		io.WriteString(w, msg)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if params.Password == "" {
		err = errors.New(missingPassword)
		msg := logger.JSONError(err)
		io.WriteString(w, msg)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if params.Name == "" {
		err = errors.New(missingPassword)
		msg := logger.JSONError(err)
		io.WriteString(w, msg)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tokenString, err := CreateJWT(params.Email)
	if err != nil {
		log.Println(err)
		err = errors.New(cantCreateJWT)
		msg := logger.JSONError(err)
		io.WriteString(w, msg)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user := &model.User{
		Email:    params.Email,
		Password: params.Password,
		Name:     params.Name,
	}

	err = db.CreateUser(user)
	if err != nil {
		err = errors.New(cantCreateUserDB)
		msg := logger.JSONError(err)
		io.WriteString(w, msg)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	rspns := &response{
		Token:   tokenString,
		Email:   user.Email,
		Message: "You are singed up successfully",
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(rspns)

}

func GetUserDetails(w http.ResponseWriter, r *http.Request) {
	log.Println("Auth: User Details")

	bearerToken := r.Header.Get("Authorization")
	authToken := strings.Split(bearerToken, " ")[0]

	email, err := VerifyToken(authToken)
	if err != nil {
		err = errors.New(cantVerifyToken)
		msg := logger.JSONError(err)
		io.WriteString(w, msg)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user, err := db.GetUserDetails(email)
	if err != nil {
		err = errors.New(cantGetDataFromDb)
		msg := logger.JSONError(err)
		io.WriteString(w, msg)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(&user)

}
