package routes

import (
	"encoding/json"
	"fmt"
	"goauth/v2/src/helpers"
	"goauth/v2/src/models"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func (h *Handler) ResetPasswordHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var user models.User
	var auth models.ResetPassword
	json.Unmarshal(body, &auth)

	result := h.DB.Where(&models.User{Email: auth.LoginId, Token: auth.Token}).Or(&models.User{Telephone: auth.LoginId}).First(&user)
	// Append to the Books table
	if result.Error != nil {
		fmt.Println(result.Error)
		response := models.ResponseObject{Message: "Invalid token or token already expired", Data: nil, Error: nil}
		helpers.Response(w).Write(response, http.StatusBadRequest)
		return
	}

	now := time.Now()
	expiryTime := user.ModifiedAt

	if now.After(expiryTime) {
		fmt.Println(result.Error)
		response := models.ResponseObject{Message: "Invalid token or token already expired", Data: nil, Error: nil}
		helpers.Response(w).Write(response, http.StatusBadRequest)
		return
	}

	password, err := helpers.HashPassword(auth.Password)
	if err != nil {
		fmt.Println(err)
		appError := models.ResponseError{Errors: err, Description: "An unexpected error occured"}
		response := models.ResponseObject{Message: appError.Description, Error: appError}
		helpers.Response(w).Write(response, http.StatusInternalServerError)
		return
	}

	user.Password = password

	user.Token = rand.Intn(999999-100000) + 100000

	h.DB.Save(&user)

	response := models.ResponseObject{Message: "Password reset successful", Data: nil, Error: nil}
	helpers.Response(w).Write(response, http.StatusOK)
}
