package main

import (
	"fmt"
	"net/http"
	"os"

	"goauth/v2/src/database"
	"goauth/v2/src/middlewares"
	"goauth/v2/src/routes"

	"github.com/gorilla/mux"
)

func main() {
	os.Setenv("SITE_TITLE", "Test Site")
	os.Setenv("DB_HOST", "db")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("DB_USERNAME", "postgres")
	os.Setenv("DB_PASSWORD", "12233e23e232e23e")
	os.Setenv("DB_NAME", "postgres")

	os.Setenv("SMTP_SERVER", "city_hotels")
	os.Setenv("SMTP_PORT", "city_hotels")
	os.Setenv("SMTP_USERNAME", "city_hotels")
	os.Setenv("SMTP_PASSWORD", "city_hotels")

	os.Setenv("AUTH_TOKEN_SECRET", "1n2kj11j2n12n4224wd35d2353wd")

	// Get the value of an Environment Variable
	host := os.Getenv("SITE_TITLE")
	port := os.Getenv("DB_HOST")

	fmt.Printf("Site status: %s RUNNING, DB RUNNING ON Host: %s\n", host, port)

	DB := database.Init()
	h := routes.New(DB)

	mainRouter := mux.NewRouter()
	authRouter := mainRouter.PathPrefix("/auth").Subrouter()

	authRouter.HandleFunc("/register", h.SignupHandler).Methods(http.MethodPost)

	authRouter.HandleFunc("/forgot-password", h.ForgotPasswordHandler).Methods(http.MethodPut)

	authRouter.HandleFunc("/reset-password", h.ResetPasswordHandler).Methods(http.MethodPut)

	authRouter.HandleFunc("/verify-token", h.VerifyTokenHandler).Methods(http.MethodPut)

	authRouter.HandleFunc("/login", h.SigninHandler).Methods(http.MethodPost)

	//Gaurds
	// Refers to a routes accessible by an authenticated user
	amw := middlewares.AuthenticationMiddleware{}
	roleGuards := middlewares.RoleMiddleware{Role: 3}

	//Role Guards
	adminRouter := mainRouter.PathPrefix("/user").Subrouter()

	adminRouter.HandleFunc("/", h.ListUserHandler).Methods(http.MethodGet)

	adminRouter.HandleFunc("/{Id}/permissions", h.ListUserHandler).Methods(http.MethodPut)

	adminRouter.HandleFunc("/{Id}/permissions", h.ListUserHandler).Methods(http.MethodGet)

	// adminRouter.HandleFunc("", h.ListUserHandler).Methods(http.MethodGet)

	adminRouter.Use(amw.Middleware, roleGuards.Middleware)

	userRouter := mainRouter.PathPrefix("/user").Subrouter()

	userRouter.HandleFunc("/change-password", h.ChangePasswordHandler).Methods(http.MethodPut)

	userRouter.HandleFunc("/profile", h.GetProfileHandler).Methods(http.MethodGet)

	userRouter.HandleFunc("/{Id}", h.GetProfileHandler).Methods(http.MethodGet)

	userRouter.Use(amw.Middleware)

	// The Signin will send the JWT back as we are making microservices.
	// The JWT token will make sure that other services are protected.
	// So, ultimately, we would need a middleware

	// mux.MiddlewareFunc()

	// Add the middleware to different subrouter
	// HTTP server
	// Add time outs
	server := &http.Server{
		Addr:    "0.0.0.0:5000",
		Handler: mainRouter,
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Error Booting the Server ", err)
	}

}
