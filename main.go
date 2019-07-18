package main

import (
	"flag"
	"net/http"
	"os"

	"github.com/Girbons/effective-potato/pkg/config"
	"github.com/Girbons/effective-potato/pkg/handlers"
	"github.com/Girbons/effective-potato/pkg/middleware"
	"github.com/gorilla/mux"
)

func main() {
	createAdminUser := flag.Bool("create-admin-user", false, "Create Admin user")
	adminPassword := flag.String("admin-password", "dev123456", "choose admin password")

	flag.Parse()

	router := mux.NewRouter()
	// db initialization
	db := config.InitDB()

	userHandlers := handlers.NewUserHandler(db)
	if *createAdminUser {
		userHandlers.CreateAdminUser(*adminPassword)
		os.Exit(2)
	}

	lightHandler := handlers.NewLightHandler(db)

	// user endpoints
	router.HandleFunc("/login/", userHandlers.Login).Methods("POST")
	router.HandleFunc("/logout/", userHandlers.Logout).Methods("POST")

	router.Handle("/api/user/add/", middleware.AuthMiddleware(http.HandlerFunc(userHandlers.AddUser))).Methods("POST")

	// lights endpoints
	router.Handle("/api/lights/list/", middleware.AuthMiddleware(http.HandlerFunc(lightHandler.List))).Methods("GET")
	router.Handle("/api/lights/add/", middleware.AuthMiddleware(http.HandlerFunc(lightHandler.Add))).Methods("POST")
	router.Handle("/api/lights/{id}/detail/", middleware.AuthMiddleware(http.HandlerFunc(lightHandler.Detail))).Methods("GET")
	router.Handle("/api/lights/{id}/detail/", middleware.AuthMiddleware(http.HandlerFunc(lightHandler.Detail))).Methods("POST")
	router.Handle("/api/lights/{id}/detail/", middleware.AuthMiddleware(http.HandlerFunc(lightHandler.Detail))).Methods("DELETE")

	// fan endpoints
	//router.Handle("/api/fan/{id}/detail/", middleware.AuthMiddleware(http.HandlerFunc(fansHandler.Detail))).Methods("DELETE")
	//router.Handle("/api/fan/list/", middleware.AuthMiddleware(http.HandlerFunc(fansHandler.List))).Methods("GET")
	//router.Handle("/api/fan/add/", middleware.AuthMiddleware(http.HandlerFunc(fansHandler.Add))).Methods("POST")
	//router.Handle("/api/fan/{id}/detail/", middleware.AuthMiddleware(http.HandlerFunc(fansHandler.Detail))).Methods("GET")
	//router.Handle("/api/fan/{id}/detail/", middleware.AuthMiddleware(http.HandlerFunc(fansHandler.Detail))).Methods("POST")
	//router.Handle("/api/fan/{id}/detail/", middleware.AuthMiddleware(http.HandlerFunc(fansHandler.Detail))).Methods("DELETE")

	//// room endpoinds
	//router.Handle("/api/room/{id}/detail/", middleware.AuthMiddleware(http.HandlerFunc(roomHandler.Detail))).Methods("DELETE")
	//router.Handle("/api/room/list/", middleware.AuthMiddleware(http.HandlerFunc(roomHandler.List))).Methods("GET")
	//router.Handle("/api/room/add/", middleware.AuthMiddleware(http.HandlerFunc(roomHandler.Add))).Methods("POST")
	//router.Handle("/api/room/{id}/detail/", middleware.AuthMiddleware(http.HandlerFunc(roomHandler.Detail))).Methods("GET")
	//router.Handle("/api/room/{id}/detail/", middleware.AuthMiddleware(http.HandlerFunc(roomHandler.Detail))).Methods("POST")
	//router.Handle("/api/room/{id}/detail/", middleware.AuthMiddleware(http.HandlerFunc(roomHandler.Detail))).Methods("DELETE")

	http.ListenAndServe(":8080", router)
}
