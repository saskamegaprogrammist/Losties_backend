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
	},
	)
	if err != nil {
		utils.WriteError(false, "Failed to create db", err)
	}

	err = useCases.Init(database.GetUsersDB())
	if err != nil {
		utils.WriteError(false, "Failed to create useCases", err)
	}

	err = losties_handlers.Init(useCases.GetUsersUC())
	if err != nil {
		utils.WriteError(false, "Failed to create handlers", err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/signup",  losties_handlers.GetUsersH().SignUp).Methods("POST")
	r.HandleFunc("/login",  losties_handlers.GetUsersH().Login).Methods("POST")


	cors := handlers.CORS(handlers.AllowedOrigins([]string{"http://localhost:3000"}), handlers.AllowCredentials())

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