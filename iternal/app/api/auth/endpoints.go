package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"oblique/iternal/app/db"
	"oblique/iternal/app/model"
	"strings"
	"time"

	"oblique/iternal/app/logger"
)

// error messages
const (
	missingEmail    = "Email is empty"
	missingPassword = "Password is empty"
	missingName     = "Name is empty"

	cantGetDataFromDb = "Couldn't get data from database"
	cantCreateUserDB  = "Couldn't create user in database"

	cantCreateJWT   = "Couldn't create JWT token"
	cantVerifyToken = "Couldn't verify token"

	cantCreateHash = "Couldn't create hash password"
)

func SignIn(w http.ResponseWriter, r *http.Request) {
	var params LoginParams

	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		msg := logger.JSONError(err)
		io.WriteString(w, msg)
		return
	}

	start := time.Now()
	valid := make(chan bool)
	go func() {
		valid <- false
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

		valid <- true
	}()
	fmt.Println("validating", time.Since(start))

	isValid := <-valid
	if isValid != true {
		return
	}

	usr := make(chan model.User)
	start = time.Now()
	go func() {
		user, err := db.GetUser(params.Email, params.Password)
		if err != nil {
			err = errors.New(cantGetDataFromDb)
			msg := logger.JSONError(err)
			io.WriteString(w, msg)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		usr <- *user
	}()

	user := <-usr
	fmt.Println("db", time.Since(start))

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
	var params RegistationParams

	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		msg := logger.JSONError(err)
		io.WriteString(w, msg)
		return
	}

	valid := make(chan bool)
	go func() {
		valid <- false

		if params.Email == "" {
			WriteError(w, missingEmail, http.StatusBadRequest)
			return
		}

		if params.Password == "" {
			WriteError(w, missingPassword, http.StatusBadRequest)
			return
		}

		if params.Name == "" {
			WriteError(w, missingName, http.StatusBadRequest)
			return
		}
		valid <- true
	}()
	if <-valid != true {
		return
	}

	hash := make(chan string)
	go func() {
		passwordHash, err := HashPassword(params.Password)
		if err != nil {
			WriteError(w, cantCreateHash, http.StatusInternalServerError)
			return
		}
		hash <- passwordHash
	}()

	user := &model.User{
		Email:    params.Email,
		Password: <-hash,
		Name:     params.Name,
	}

	errChan := make(chan error)
	go func() {
		err = db.CreateUser(user)
		errChan <- err
	}()

	if <-errChan != nil {
		WriteError(w, cantCreateUserDB, http.StatusInternalServerError)
		return
	}

	token := make(chan string)
	go func() {
		tokenString, err := CreateJWT(params.Email)
		if err != nil {
			log.Println(err)
			WriteError(w, cantCreateJWT, http.StatusInternalServerError)
			return
		}
		token <- tokenString
	}()

	rspns := &response{
		Token:   <-token,
		Email:   user.Email,
		Message: "You are singed up successfully",
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rspns)
}

func GetUserDetails(w http.ResponseWriter, r *http.Request) {
	bearerToken := r.Header.Get("Authorization")
	authToken := strings.Split(bearerToken, " ")[0]

	email, err := VerifyToken(authToken)
	if err != nil {
		WriteError(w, cantVerifyToken, http.StatusInternalServerError)
		return
	}

	user, err := db.GetUserDetails(email)
	if err != nil {
		WriteError(w, cantVerifyToken, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(&user)
}
