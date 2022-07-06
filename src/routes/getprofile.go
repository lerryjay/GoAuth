package routes

import (
	"encoding/json"
	"goauth/v2/src/helpers"
	"goauth/v2/src/models"
	"net/http"

	"github.com/gorilla/context"
)

// TODO: Implement get user profile

func (h *Handler) GetProfileHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	data := context.Get(r, "user")
	json.Unmarshal([]byte(data.(string)), &user)

	response := models.ResponseObject{Message: "Profile retrieved successfully", Data: user, Error: nil}
	helpers.Response(w).Write(response, http.StatusOK)
}
