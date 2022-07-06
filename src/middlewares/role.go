package middlewares

import (
	"encoding/json"
	"goauth/v2/src/models"
	"log"
	"net/http"

	"github.com/gorilla/context"
)

// Define our struct
type RoleMiddleware struct {
	Role int
}

// Middleware function, which will be called for each request
func (amw *RoleMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var user models.User
		data := context.Get(r, "user")
		json.Unmarshal([]byte(data.(string)), &user)

		if user.Role == amw.Role {
			// user, err := jwt.decode(token, os.Getenv("AUTH_TOKEN_SECRET")
			// var userM models.User
			// json.Unmarshal(user, &userM)
			// fmt.Print(user, userM)

			// We found the token in our map
			// Pass down the request to the next middleware (or final handler)

			next.ServeHTTP(w, r)
		} else {
			log.Println("User does not have permission")
			// Write an error and stop the handler chain
			http.Error(w, "Forbidden", http.StatusForbidden)
		}
	})
}
