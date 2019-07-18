package handlers

import (
	"encoding/json"
	"fmt"
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
	params := r.URL.Query()

	name := params.Get("name")
	pin := params.Get("pin")

	if name == "" || pin == "" {
		respondWithError(w, http.StatusBadRequest, "please add pin or name")
	} else {
		i, _ := strconv.Atoi(pin)
		l.db.Where(&models.Light{Name: name, Pin: i}).FirstOrCreate(&models.Light{})
		respondWithJSON(w, http.StatusOK, fmt.Sprintf("light %s with pin %s correctly saved", name, pin))
	}
}

func (l *LightHandler) Detail(w http.ResponseWriter, r *http.Request) {
	var light models.Light

	vars := mux.Vars(r)
	id, _ := strconv.ParseUint(vars["id"], 10, 32)

	l.db.First(&light, int(id)).Scan(&light)
	fmt.Println(light.ID)

	if light.ID == 0 {
		respondWithError(w, http.StatusBadRequest, "cannot find light with id "+vars["id"])
	}

	switch r.Method {
	case "POST":
		params := r.URL.Query()

		name := params.Get("name")
		pin := params.Get("pin")

		if name != "" {
			light.Name = name
			l.db.Save(&light)

		}
		if pin != "" {
			i, _ := strconv.Atoi(pin)
			light.Pin = i
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
