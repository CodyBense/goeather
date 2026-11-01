package main

import (
	"net/http"

	"github.com/CodyBense/goeather/internal/api"
	"github.com/CodyBense/goeather/internal/json"
)

// handlerWeatherGet ...
func (cfg *apiConfig) handlerWeatherGet(w http.ResponseWriter, r *http.Request) {
	weather, err := api.GetWeather(cfg.apiKey)
	if err != nil {
		json.RespondWithError(w, http.StatusInternalServerError, "Couldn't get weather", err)
	}

	json.RespondWithJson(w, http.StatusOK, weather)
}
