package controllers

import (
	"encoding/json"
	"fmt"
	"golang-web/src/helper"
	"golang-web/src/models"
	"net/http"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var input models.User
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "invalid request body")
			return
		}
		hashPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
		Password := string(hashPassword)
		newUser := models.User{
			Email:    input.Email,
			Password: Password,
		}
		res := models.CreateUser(&newUser)
		var _, _ = json.Marshal(res)
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintln(w, "Register Succesful")
	} else {
		http.Error(w, "", http.StatusBadRequest)
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var input models.User
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "invalid request body")
			return
		}
		ValidateEmail := models.FindEmail(&input)
		if len(ValidateEmail) == 0 {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintln(w, "Email is not Found")
			return
		}
		var passwordSecond string
		for _, user := range ValidateEmail {
			passwordSecond = user.Password
		}
		if err := bcrypt.CompareHashAndPassword([]byte(passwordSecond), []byte(input.Password)); err != nil {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "Password not Found")
			return
		}
		jwtKey := os.Getenv("SECRETKEY")
		token, _ := helper.GenerateToken(jwtKey, input.Email)
		item := map[string]string{
			"Email": input.Email,
			"Token": token,
		}
		res, _ := json.Marshal(item)
		w.WriteHeader(http.StatusOK)
		w.Write(res)
		return
	} else {
		http.Error(w, "", http.StatusBadRequest)
	}
}
