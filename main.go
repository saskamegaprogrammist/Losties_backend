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
		utils.WriteError(false, "Failed	 to create db", err)
	}

	err = useCases.Init(database.GetUsersDB(), database.GetCookiesDB(), database.GetFilenameDB(), database.GetAdsDB(),  database.GetPetsDB(), database.GetCoordsDB(), database.GetCommentsDB())
	if err != nil {
		utils.WriteError(false, "Failed to create useCases", err)
	}

	err = losties_handlers.Init(useCases.GetUsersUC(), useCases.GetAdsUC(), useCases.GetPetsUC(), useCases.GetPicUC(), useCases.GetCoordsUC(), useCases.GetCommentsUC())
	if err != nil {
		utils.WriteError(false, "Failed to create handlers", err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/auth",  losties_handlers.GetUsersH().Auth).Methods("GET")
	r.HandleFunc("/logout",  losties_handlers.GetUsersH().Logout).Methods("DELETE")
	r.HandleFunc("/signup",  losties_handlers.GetUsersH().SignUp).Methods("POST")
	r.HandleFunc("/login",  losties_handlers.GetUsersH().Login).Methods("POST")
	r.HandleFunc("/user/{id}",  losties_handlers.GetUsersH().Update).Methods("PUT")
	r.HandleFunc("/user/{id}",  losties_handlers.GetUsersH().GetUser).Methods("GET")
	r.HandleFunc("/user/{id}/card",  losties_handlers.GetAdsH().NewAd).Methods("POST")
	r.HandleFunc("/user/{id}/cards",  losties_handlers.GetAdsH().GetUserAds).Methods("GET")
	r.HandleFunc("/user/{id}/cards/number",  losties_handlers.GetAdsH().GetUserAdsNumber).Methods("GET")
	r.HandleFunc("/card/{id}",  losties_handlers.GetAdsH().GetAd).Methods("GET")
	r.HandleFunc("/card/{id}/pet",  losties_handlers.GetPetsH().NewPet).Methods("POST")
	r.HandleFunc("/card/{id}/pet",  losties_handlers.GetPetsH().GetAdPet).Methods("GET")
	r.HandleFunc("/cards",  losties_handlers.GetAdsH().GetAds).Methods("GET")
	r.HandleFunc("/card/{id}/pet",  losties_handlers.GetPetsH().NewPet).Methods("POST")
	r.HandleFunc("/card/{id}/pic",  losties_handlers.GetPicH().NewAdPic).Methods("POST")
	r.HandleFunc("/card/{id}/pic",  losties_handlers.GetPicH().AdPicGet).Methods("GET")
	r.HandleFunc("/user/{id}/pic",  losties_handlers.GetPicH().NewUserPic).Methods("POST")
	r.HandleFunc("/user/{id}/pic",  losties_handlers.GetPicH().UserPicGet).Methods("GET")
	r.HandleFunc("/card/{id}/coords",  losties_handlers.GetCoordsH().NewCoords).Methods("POST")
	r.HandleFunc("/card/{id}/coords",  losties_handlers.GetCoordsH().GetAdCoords).Methods("GET")
	r.HandleFunc("/coords",  losties_handlers.GetCoordsH().GetCoords).Methods("GET")
	r.HandleFunc("/card/{id}/comments",  losties_handlers.GetCommentsH().GetAdComments).Methods("GET")
	r.HandleFunc("/card/{id}/comment",  losties_handlers.GetCommentsH().NewComment).Methods("POST")


	cors := handlers.CORS(handlers.AllowedOrigins([]string{"http://localhost:5000", "warm-chamber-57831.herokuapp.com", "http://warm-chamber-57831.herokuapp.com", "https://warm-chamber-57831.herokuapp.com"}), handlers.AllowCredentials(), handlers.AllowedMethods([]string{"POST", "GET", "PUT", "DELETE"}))

	server := &http.Server{
		Addr: ":5001",
		Handler : cors(r),
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	err = server.ListenAndServe()
	if err != nil {
		logger.Errorf("Failed to start server: %v", err)
	}
}