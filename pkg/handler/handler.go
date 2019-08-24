package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Girbons/effective-potato/pkg/device"
	"github.com/Girbons/effective-potato/pkg/sensor"
	log "github.com/sirupsen/logrus"
)

func PinON(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()

	pin := params.Get("pin")

	if pin == "" {
		respondWithError(w, http.StatusBadRequest, "Pin is required")
	}

	parsedPin, _ := strconv.Atoi(pin)
	device.On(parsedPin)

	respondWithJSON(w, http.StatusOK, fmt.Sprintf("pin set to on"))
}

func PinOFF(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()

	pin := params.Get("pin")

	if pin == "" {
		respondWithError(w, http.StatusBadRequest, "Pin is required")
	}

	parsedPin, _ := strconv.Atoi(pin)
	device.Off(parsedPin)

	respondWithJSON(w, http.StatusOK, fmt.Sprintf("pin set to off"))
}

func PinStatus(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()

	pin := params.Get("pin")

	if pin == "" {
		respondWithError(w, http.StatusBadRequest, "Pin is required")
	}

	parsedPin, _ := strconv.Atoi(pin)
	status := device.Status(parsedPin)

	respondWithJSON(w, http.StatusOK, fmt.Sprintf("status %s", status))
}

func ReadTemperature(w http.ResponseWriter, r *http.Request) {
	retryTimes := 1

	params := r.URL.Query()

	pin := params.Get("pin")
	dhtType := params.Get("dht")
	retryTimesParam := params.Get("retryTimes")

	if pin == "" || dhtType == "" {
		respondWithError(w, http.StatusBadRequest, "Pin and DHT are required")
	}

	if retryTimesParam != "" {
		i, _ := strconv.Atoi(retryTimesParam)
		retryTimes = i
	}

	parsedPin, _ := strconv.Atoi(pin)
	temperature, humidity, _, err := sensor.ReadTemperature(parsedPin, dhtType, retryTimes)

	if err != nil {
		log.Error(err)
	}

	resp := make(map[string]float32)
	resp["temperature"] = temperature
	resp["humidity"] = humidity

	respondWithJSON(w, http.StatusOK, resp)
}
