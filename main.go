package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx"
	"github.com/saskamegaprogrammist/Losties_backend/database"
	"os"
	"time"

	"github.com/google/logger"
	"net/http"
)

const logPath  = "log.log"
const verbose = true

func loggerSetup() {
	lf, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
	if err != nil {
		logger.Fatalf("Failed to open log file: %v", err)
	}
	defer lf.Close()
	defer logger.Init("LoggerExample", verbose, true, lf).Close()
}

func main() {

	loggerSetup()

	err := database.Init(pgx.ConnConfig{
		Host:     "localhost",
		User:     "alexis",
		Password: "sinope27",
	},
	)
	if err != nil {
		logger.Errorf("Failed to create db: %v", err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/signup", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "text/plain")
		writer.Write([]byte("This is an example server.\n"))
	})

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