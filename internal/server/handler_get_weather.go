package main

import (
	"net/http"

	"github.com/CodyBense/goeather/internal/api"
)

// handlerWeatherGet ...
func (cfg *apiConfig) handlerWeatherGet(w http.ResponseWriter, r *http.Request) {
	weather, err := api.GetWeather(cfg.apiKey)
	respondWithError(w, http.StatusInternalServerError, "Couldn't get weather")

	respondWithJson(w, http.StatusOK, weather)
}
