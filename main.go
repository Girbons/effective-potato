package main

import (
	"flag"
	"net/http"
	"os"

	"github.com/Girbons/effective-potato/pkg/config"
	"github.com/Girbons/effective-potato/pkg/handler"
	"github.com/Girbons/effective-potato/pkg/middleware"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/stianeikeland/go-rpio"
)

func main() {
	// used to create the admin user
	createAdminUser := flag.Bool("create-admin-user", false, "Create Admin user")
	adminPassword := flag.String("admin-password", "123456", "choose admin password")
	flag.Parse()

	// initialize database
	db := config.InitDB()
	defer db.Close()

	userHandlers := handler.NewUserHandler(db)
	if *createAdminUser {
		userHandlers.CreateAdminUser(*adminPassword)
		os.Exit(0)
	}

	if err := rpio.Open(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
	defer rpio.Close()

	router := mux.NewRouter()
	// user endpoints
	router.HandleFunc("/login/", userHandlers.Login).Methods("POST")
	router.HandleFunc("/logout/", userHandlers.Logout).Methods("POST")
	router.Handle("/user/profile/", middleware.AuthMiddleware(http.HandlerFunc(userHandlers.Profile))).Methods("GET")

	router.Handle("/api/pin/on/{pin:[0-9]+}/", middleware.AuthMiddleware(http.HandlerFunc(handler.PinON))).Methods("GET")
	router.Handle("/api/pin/off/{pin:[0-9]+}/", middleware.AuthMiddleware(http.HandlerFunc(handler.PinOFF))).Methods("GET")
	router.Handle("/api/pin/status/{pin:[0-9]+}/", middleware.AuthMiddleware(http.HandlerFunc(handler.PinStatus))).Methods("GET")
	router.Handle("/api/temperature-sensor/{pin:[0-9]+}/{dht}/", middleware.AuthMiddleware(http.HandlerFunc(handler.ReadTemperature))).Methods("GET")

	http.ListenAndServe(":8080", router)
}
