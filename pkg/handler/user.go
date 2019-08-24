package handler

import (
	"fmt"
	"net/http"

	"github.com/Girbons/effective-potato/pkg/auth"
	"github.com/Girbons/effective-potato/pkg/models"
	"github.com/bitly/go-simplejson"
	"github.com/jinzhu/gorm"
)

type UserHandler struct {
	db *gorm.DB
}

func NewUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{
		db: db,
	}
}

func (user *UserHandler) CreateAdminUser(password string) {
	var u models.User

	hash, _ := auth.HashPassword(password)

	user.db.Where(&models.User{Username: "admin", Password: hash}).FirstOrCreate(&u)
	fmt.Println("Admin user correctly created")
}

// AddUserHandler takes care of creating a new user given `username` and `password`
func (user *UserHandler) AddUser(w http.ResponseWriter, r *http.Request) {
	var u models.User

	params := r.URL.Query()

	username := params.Get("username")
	password := params.Get("password")

	if username == "" || password == "" {
		respondWithError(w, http.StatusBadRequest, "provide username and password")
	} else {

		// here we hash the password and we drop it
		// we won't save the password in plain text OFC
		hash, _ := auth.HashPassword(password)

		user.db.Where(&models.User{Username: "admin", Password: hash}).FirstOrCreate(&u)
		respondWithJSON(w, http.StatusCreated, "user correctly created")
	}
}

// Login takes care to check if a User can login
func (user *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var u models.User
	var count int

	params := r.URL.Query()

	username := params.Get("username")
	password := params.Get("password")

	if username == "" || password == "" {
		respondWithError(w, http.StatusBadRequest, "provide username and password")
	}

	user.db.Where("username = ?", username).First(&u).Count(&count)
	fmt.Println("count", count)
	if count == 0 {
		respondWithError(w, http.StatusBadRequest, "invalid username or password")
	}

	hash, _ := auth.HashPassword(password)
	if ok := auth.CheckPasswordHash(password, hash); !ok {
		respondWithError(w, http.StatusBadRequest, "invalid username or password")
	}

	token, err := auth.GetToken(username)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "There was an error trying to generate JWT "+err.Error())
	}

	json := simplejson.New()
	json.Set("token", token)

	payload, err := json.MarshalJSON()
	if err != nil {
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Authorization", "Bearer "+token)
	w.Write(payload)
}

// Logout endpoint is used to log a use out
func (user *UserHandler) Logout(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, "Logged Out succesfully")
}
