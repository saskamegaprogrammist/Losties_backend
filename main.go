package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx"
	"github.com/saskamegaprogrammist/Losties_backend/database"
	losties_handlers "github.com/saskamegaprogrammist/Losties_backend/handlers"
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

	r := mux.NewRouter()
	r.HandleFunc("/signup",  losties_handlers.SignUp).Methods("POST")

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