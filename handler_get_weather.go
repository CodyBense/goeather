package main

import (
	"encoding/json"
	"net/http"

	"github.com/CodyBense/goeather/internal/api"
	"github.com/CodyBense/goeather/internal/cjson"
)

// handlerWeatherGetHourl ...
func (cfg *apiConfig) handlerWeatherGet(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		ApiKey string `json:"api_key"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		cjson.RespondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
	}

	weather, err := api.GetWeather(params.ApiKey)
	if err != nil {
		cjson.RespondWithError(w, http.StatusInternalServerError, "Couldn't get hourly weather", err)
	}

	cjson.RespondWithJson(w, http.StatusOK, weather)
}
