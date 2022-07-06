package middlewares

import (
	"goauth/v2/src/helpers"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/context"
)

// Define our struct
type AuthenticationMiddleware struct {
}

// Middleware function, which will be called for each request
func (amw *AuthenticationMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		token = strings.Split(token, "Bearer ")[1]

		if user, err := helpers.ValidateToken(token, os.Getenv("AUTH_TOKEN_SECRET")); user != nil {
			// user, err := jwt.decode(token, os.Getenv("AUTH_TOKEN_SECRET")
			// var userM models.User
			// json.Unmarshal(user, &userM)
			// fmt.Print(user, userM)

			// We found the token in our map
			// Pass down the request to the next middleware (or final handler)
			context.Set(r, "user", user)

			next.ServeHTTP(w, r)
		} else {
			if err != nil {
				log.Println(err, token)
			}
			// Write an error and stop the handler chain
			http.Error(w, "Forbidden", http.StatusForbidden)
		}
	})
}
