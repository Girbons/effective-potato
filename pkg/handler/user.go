package handler

import (
	"encoding/json"
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

func (user *UserHandler) Profile(w http.ResponseWriter, r *http.Request) {
	var u models.User

	username := r.Header.Get("username")
	user.db.Where("username = ?", username).First(&u)

	result, _ := json.Marshal(u)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

type IntUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Login takes care to check if a User can login
func (user *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var u models.User
	var us IntUser
	var count int

	json.NewDecoder(r.Body).Decode(&us)

	user.db.Where("username = ?", us.Username).First(&u).Count(&count)
	if count == 0 {
		respondWithError(w, http.StatusBadRequest, "invalid username or password")
	}

	hash, _ := auth.HashPassword(us.Password)
	if ok := auth.CheckPasswordHash(us.Password, hash); !ok {
		respondWithError(w, http.StatusBadRequest, "invalid username or password")
	}

	token, err := auth.GetToken(us.Username)
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
