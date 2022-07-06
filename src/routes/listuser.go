package routes

import (
	"goauth/v2/src/helpers"
	"goauth/v2/src/models"
	"log"
	"net/http"
)

// TODO: Implement get user profile

func (h *Handler) ListUserHandler(w http.ResponseWriter, r *http.Request) {

	var users []models.User

	result := h.DB.Find(&users)
	// Append to the Books table
	if result.Error != nil {
		log.Println(result.Error)
		response := models.ResponseObject{Message: "Internal Server Error"}
		helpers.Response(w).Write(response, http.StatusInternalServerError)
		return
	}

	response := models.ResponseObject{Message: "Fetched users successfully", Data: users, Error: nil}
	helpers.Response(w).Write(response, http.StatusOK)
}
