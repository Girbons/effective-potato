package handler

import (
	"net/http"
	"strconv"

	"github.com/Girbons/effective-potato/pkg/device"
	"github.com/Girbons/effective-potato/pkg/sensor"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func PinON(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pin := vars["pin"]

	if pin == "" {
		respondWithError(w, http.StatusBadRequest, "Pin is required")
	}

	parsedPin, _ := strconv.Atoi(pin)
	device.On(parsedPin)

	respondWithJSON(w, http.StatusOK, map[string]string{"status": "pin ON"})
}

func PinOFF(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pin := vars["pin"]

	if pin == "" {
		respondWithError(w, http.StatusBadRequest, "Pin is required")
	}

	parsedPin, _ := strconv.Atoi(pin)
	device.Off(parsedPin)

	respondWithJSON(w, http.StatusOK, map[string]string{"status": "pin OFF"})
}

func PinStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pin := vars["pin"]

	if pin == "" {
		respondWithError(w, http.StatusBadRequest, "Pin is required")
	}

	parsedPin, _ := strconv.Atoi(pin)
	status := device.Status(parsedPin)

	respondWithJSON(w, http.StatusOK, map[string]string{"status": status})
}

func ReadTemperature(w http.ResponseWriter, r *http.Request) {
	retryTimes := 1

	params := r.URL.Query()
	vars := mux.Vars(r)
	pin := vars["pin"]

	dhtType := vars["dht"]
	retryTimesParam := params.Get("retryTimes")

	if pin == "" || dhtType == "" {
		respondWithError(w, http.StatusBadRequest, "Pin and DHT are required")
	}

	if retryTimesParam != "" {
		i, _ := strconv.Atoi(retryTimesParam)
		retryTimes = i
	}

	parsedPin, _ := strconv.Atoi(pin)
	temperature, humidity, _, err := sensor.ReadTemperature(dhtType, parsedPin, retryTimes)

	if err != nil {
		log.Error(err)
	}

	resp := make(map[string]float32)
	resp["temperature"] = temperature
	resp["humidity"] = humidity

	respondWithJSON(w, http.StatusOK, resp)
}
