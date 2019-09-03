package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Girbons/effective-potato/pkg/models"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type LightHandler struct {
	db *gorm.DB
}

func NewLightHandler(db *gorm.DB) *LightHandler {
	return &LightHandler{
		db: db,
	}
}

func (l *LightHandler) Add(w http.ResponseWriter, r *http.Request) {
	var light models.Light

	json.NewDecoder(r.Body).Decode(&light)

	if light.Name == "" || light.Pin == nil {
		respondWithError(w, http.StatusBadRequest, "please add pin or name")
	} else {
		l.db.Where(&models.Light{Name: light.Name, Pin: light.Pin}).FirstOrCreate(&models.Light{})
		respondWithJSON(w, http.StatusOK, "light created")
	}
}

func (l *LightHandler) Detail(w http.ResponseWriter, r *http.Request) {
	var light models.Light

	vars := mux.Vars(r)
	id, _ := strconv.ParseUint(vars["id"], 10, 32)

	l.db.First(&light, int(id)).Scan(&light)

	if light.ID == 0 {
		respondWithError(w, http.StatusBadRequest, "cannot find light with id "+vars["id"])
	}

	switch r.Method {
	case "POST":
		var newLight models.Light

		json.NewDecoder(r.Body).Decode(&newLight)

		if newLight.Name != "" {
			light.Name = newLight.Name
			l.db.Save(&light)

		}
		if newLight.Pin != nil {
			light.Pin = newLight.Pin
			l.db.Save(&light)
		}
		break
	case "DELETE":
		l.db.Delete(&light)
		respondWithJSON(w, http.StatusNoContent, "")
		break
	default:
		result, _ := json.Marshal(light)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(result)

	}

}

func (l *LightHandler) List(w http.ResponseWriter, r *http.Request) {
	var lights []*models.Light

	l.db.Find(&lights).Scan(&lights)

	results, _ := json.Marshal(lights)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(results)
}
