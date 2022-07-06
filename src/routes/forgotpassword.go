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
	"os"
	"strconv"
)

func (h *Handler) ForgotPasswordHandler(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var user models.User
	var auth models.Login
	json.Unmarshal(body, &auth)

	result := h.DB.Where(&models.User{Email: auth.LoginId}).Or(&models.User{Telephone: auth.LoginId}).First(&user)
	// Append to the Books table
	if result.Error != nil {
		fmt.Println(result.Error)
		response := models.ResponseObject{Message: "An email has been sent to you. Please follow the intructions to reset your password", Data: nil, Error: nil}
		helpers.Response(w).Write(response, http.StatusOK)
		return
	}

	user.Token = rand.Intn(999999-100000) + 100000

	h.DB.Save(&user)

	mail := fmt.Sprintf("<p>A request was made to reset your password. Please use the following token to continue your password reest: <br /><br /> <b>%s</b> </p>", strconv.Itoa(user.Token))
	subject := fmt.Sprintf("%s :Reset Password", os.Getenv("SITE_TITLE"))
	helpers.SendMail(user.Email, subject, mail)

	response := models.ResponseObject{Message: "An email has been sent to you. Please follow the intructions to reset your password", Data: nil, Error: nil}
	helpers.Response(w).Write(response, http.StatusOK)
}
