package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx"
	"github.com/saskamegaprogrammist/Losties_backend/database"
	losties_handlers "github.com/saskamegaprogrammist/Losties_backend/handlers"
	"github.com/saskamegaprogrammist/Losties_backend/useCases"
	"github.com/saskamegaprogrammist/Losties_backend/utils"
	"time"

	"github.com/google/logger"
	"net/http"
)



func main() {
	utils.LoggerSetup()
	defer utils.Close()

	err := database.Init(pgx.ConnConfig{
		Database: "losties",
		Host:     "localhost",
		User:     "alexis",
		Password: "sinope27",
	}, "localhost")
	if err != nil {
		utils.WriteError(false, "Failed to create db", err)
	}

	err = useCases.Init(database.GetUsersDB(), database.GetCookiesDB())
	if err != nil {
		utils.WriteError(false, "Failed to create useCases", err)
	}

	err = losties_handlers.Init(useCases.GetUsersUC())
	if err != nil {
		utils.WriteError(false, "Failed to create handlers", err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/auth",  losties_handlers.GetUsersH().Auth).Methods("GET")
	r.HandleFunc("/signup",  losties_handlers.GetUsersH().SignUp).Methods("POST")
	r.HandleFunc("/login",  losties_handlers.GetUsersH().Login).Methods("POST")
	r.HandleFunc("/user/{id}",  losties_handlers.GetUsersH().Update).Methods("PUT")


	cors := handlers.CORS(handlers.AllowedOrigins([]string{"http://localhost:3000"}), handlers.AllowCredentials(), handlers.AllowedMethods([]string{"POST", "GET", "PUT", "DELETE"}))

	server := &http.Server{
		Addr: ":5000",
		Handler : cors(r),
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	err = server.ListenAndServe()
	if err != nil {
		logger.Errorf("Failed to start server: %v", err)
	}
}