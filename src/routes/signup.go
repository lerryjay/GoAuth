package routes

import (
	"encoding/json"
	"fmt"
	"goauth/v2/src/helpers"
	"goauth/v2/src/models"
	"io/ioutil"
	"log"
	"net/http"
)

type ValidationError struct {
	Field string
	Error string
}

func validateUserRegister(user models.User) (bool, []ValidationError) {
	var errors []ValidationError = []ValidationError{}
	var error bool
	if user.Firstname == "" {
		error = true
		errors = append(errors, ValidationError{Field: "Firstname", Error: "Firstname cannot be empty"})
	}

	if user.Lastname == "" {
		error = true
		errors = append(errors, ValidationError{Field: "Lastname", Error: "Lastname cannot be empty"})
	}

	if user.Email == "" {
		error = true
		errors = append(errors, ValidationError{Field: "Email", Error: "Email cannot be empty"})
	}

	if user.Telephone == "" {
		error = true
		errors = append(errors, ValidationError{Field: "Telephone", Error: "Telephone cannot be empty"})
	}

	if user.Password == "" {
		error = true
		errors = append(errors, ValidationError{Field: "Password", Error: "Password cannot be empty"})
	}
	return error, errors
}

// adds the user to the database of users
func (h *Handler) SignupHandler(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var user models.User
	json.Unmarshal(body, &user)

	invalid, errors := validateUserRegister(user)

	if invalid {
		fmt.Println(errors)
		validationError := models.ResponseError{Errors: errors, Description: "One or more validation error has occured"}
		response := models.ResponseObject{Message: validationError.Description, Error: validationError}
		helpers.Response(w).Write(response, http.StatusBadRequest)
		return
	}

	password, err := helpers.HashPassword(user.Password)
	if err != nil {
		fmt.Println(err)
		appError := models.ResponseError{Errors: err, Description: "One or more validation error has occured"}
		response := models.ResponseObject{Message: appError.Description, Error: appError}
		helpers.Response(w).Write(response, http.StatusInternalServerError)
		return
	}

	user.Password = password
	// Append to the Books table
	if result := h.DB.Create(&user); result.Error != nil {
		fmt.Println(result.Error)
	}

	tokenString, err := getSignedToken(user)
	if err != nil {
		fmt.Println(err)
		response := models.ResponseObject{Message: "Internal Server Error", Data: err}
		helpers.Response(w).Write(response, http.StatusInternalServerError)
		return
	}

	// Send a 201 created response
	response := models.ResponseObject{Message: "Registration Successful", Data: tokenString, Error: nil}
	helpers.Response(w).Write(response, http.StatusCreated)
	return
}
