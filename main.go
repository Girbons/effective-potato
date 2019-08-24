package main

import (
	"flag"
	"net/http"
	"os"

	"github.com/Girbons/effective-potato/pkg/config"
	"github.com/Girbons/effective-potato/pkg/handler"
	"github.com/Girbons/effective-potato/pkg/middleware"
	"github.com/gorilla/mux"
)

func main() {
	createAdminUser := flag.Bool("create-admin-user", false, "Create Admin user")
	adminPassword := flag.String("admin-password", "dev123456", "choose admin password")
	flag.Parse()

	db := config.InitDB()
	defer db.Close()

	userHandlers := handler.NewUserHandler(db)
	if *createAdminUser {
		userHandlers.CreateAdminUser(*adminPassword)
		os.Exit(0)
	}

	router := mux.NewRouter()
	// user endpoints
	router.HandleFunc("/login/", userHandlers.Login).Methods("POST")
	router.HandleFunc("/logout/", userHandlers.Logout).Methods("POST")

	router.Handle("/api/user/add/", middleware.AuthMiddleware(http.HandlerFunc(userHandlers.AddUser))).Methods("POST")

	router.Handle("/api/pin/on/{pin:[0-9]+}/", middleware.AuthMiddleware(http.HandlerFunc(handler.PinON))).Methods("POST")
	router.Handle("/api/pin/off/{pin:[0-9]+}/", middleware.AuthMiddleware(http.HandlerFunc(handler.PinOFF))).Methods("POST")
	router.Handle("/api/pin/status/{pin:[0-9]+}/", middleware.AuthMiddleware(http.HandlerFunc(handler.PinStatus))).Methods("GET")
	router.Handle("/api/temperature-sensor/{pin:[0-9]+}/", middleware.AuthMiddleware(http.HandlerFunc(handler.ReadTemperature))).Methods("GET")

	http.ListenAndServe(":8080", router)
}
