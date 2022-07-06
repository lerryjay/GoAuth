package routes

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"encoding/json"
	"io/ioutil"

	"goauth/v2/src/helpers"
	"goauth/v2/src/jwt"
	"goauth/v2/src/models"
)

// we need this function to be private
func getSignedToken(user models.User) (string, error) {

	userJSON, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	// we make a JWT Token here with signing method of ES256 and claims.
	// claims are attributes.
	// aud - audience
	// iss - issuer
	// exp - expiration of the Token
	claimsMap := map[string]string{
		"aud":  "frontend.knowsearch.ml",
		"iss":  "knowsearch.ml",
		"exp":  fmt.Sprint(time.Now().Add(time.Minute * 1).Unix()),
		"user": string(userJSON),
	}
	// here we provide the shared secret. It should be very complex.
	// Also, it should be passed as a System Environment variable

	secret := os.Getenv("AUTH_TOKEN_SECRET")
	header := "HS256"
	tokenString, err := jwt.GenerateToken(header, claimsMap, secret)
	if err != nil {
		return tokenString, err
	}
	return tokenString, nil
}

func (h *Handler) SigninHandler(w http.ResponseWriter, r *http.Request) {

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
		response := models.ResponseObject{Message: "Invalid username/password", Data: err}
		helpers.Response(w).Write(response, http.StatusBadRequest)
		return
	}

	valid := helpers.ValidatePasswordHash(user.Password, auth.Password)
	fmt.Print("Data", valid)

	if !valid {
		fmt.Println(err)
		response := models.ResponseObject{Message: "Invalid username/password", Data: err}
		helpers.Response(w).Write(response, http.StatusBadRequest)
		return
	}

	tokenString, err := getSignedToken(user)
	if err != nil {
		fmt.Println(err)
		response := models.ResponseObject{Message: "Internal Server Error", Data: err}
		helpers.Response(w).Write(response, http.StatusInternalServerError)
		return
	}

	response := models.ResponseObject{Message: "Login Successful", Data: tokenString, Error: nil}
	helpers.Response(w).Write(response, http.StatusOK)
}
