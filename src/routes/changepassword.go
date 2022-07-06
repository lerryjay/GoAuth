package routes

import (
	"encoding/json"
	"fmt"
	"goauth/v2/src/helpers"
	"goauth/v2/src/models"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/context"
)

func (h *Handler) ChangePasswordHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var auth models.ChangePassword
	json.Unmarshal(body, &auth)

	var user models.User
	data := context.Get(r, "user")
	json.Unmarshal([]byte(data.(string)), &user)

	if valid := helpers.ValidatePasswordHash(user.Password, auth.OldPassword); !valid {
		fmt.Println(err)
		response := models.ResponseObject{Message: "Invalid old authentication password", Data: err}
		helpers.Response(w).Write(response, http.StatusBadRequest)
		return
	}

	password, err := helpers.HashPassword(auth.NewPassword)
	if err != nil {
		fmt.Println(err)
		appError := models.ResponseError{Errors: err, Description: "An unexpected error occured"}
		response := models.ResponseObject{Message: appError.Description, Error: appError}
		helpers.Response(w).Write(response, http.StatusInternalServerError)
		return
	}

	result := h.DB.Model(&user).Where(&models.User{Email: user.Email}).Update("Password", password)
	// Append to the Books table
	if result.Error != nil {
		fmt.Println(result.Error)
		response := models.ResponseObject{Message: "Invalid username/password", Data: err}
		helpers.Response(w).Write(response, http.StatusBadRequest)
		return
	}

	response := models.ResponseObject{Message: "Password reset successful", Data: nil, Error: nil}
	helpers.Response(w).Write(response, http.StatusOK)
}
