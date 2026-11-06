package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type WeatherResp struct {
	Lat       string `json:"lat"`
	Lon       string `json:"lon"`
	Elevation int    `json:"elevation"`
	Timezone  string `json:"timezone"`
	Units     string `json:"units"`
	Current   struct {
		Icon        string  `json:"icon"`
		IconNum     int     `json:"icon_num"`
		Summary     string  `json:"summary"`
		Temperature float64 `json:"temperature"`
		Wind        struct {
			Speed float64 `json:"speed"`
			Angle int     `json:"angle"`
			Dir   string  `json:"dir"`
		} `json:"wind"`
		Precipitation struct {
			Total float64 `json:"total"`
			Type  string  `json:"type"`
		} `json:"precipitation"`
		CloudCover int `json:"cloud_cover"`
	} `json:"current"`
	Hourly struct {
		Data []struct {
			Date        string  `json:"date"`
			Weather     string  `json:"weather"`
			Icon        int     `json:"icon"`
			Summary     string  `json:"summary"`
			Temperature float64 `json:"temperature"`
			Wind        struct {
				Speed float64 `json:"speed"`
				Dir   string  `json:"dir"`
				Angle int     `json:"angle"`
			} `json:"wind"`
			CloudCover struct {
				Total int `json:"total"`
			} `json:"cloud_cover"`
			Precipitation struct {
				Total float64 `json:"total"`
				Type  string  `json:"type"`
			} `json:"precipitation"`
		} `json:"data"`
	} `json:"hourly"`
	Daily interface{} `json:"daily"`
}

type WeatherRespCurrent struct {
	Icon        string  `json:"icon"`
	IconNum     int     `json:"icon_num"`
	Summary     string  `json:"summary"`
	Temperature float64 `json:"temperature"`
	Wind        struct {
		Speed float64 `json:"speed"`
		Angle int     `json:"angle"`
		Dir   string  `json:"dir"`
	} `json:"wind"`
	Precipitation struct {
		Total float64 `json:"total"`
		Type  string  `json:"type"`
	} `json:"precipitation"`
	CloudCover int `json:"cloud_cover"`
}

type WeatherRespHourly struct {
	Data []struct {
		Date        string  `json:"date"`
		Weather     string  `json:"weather"`
		Icon        int     `json:"icon"`
		Summary     string  `json:"summary"`
		Temperature float64 `json:"temperature"`
		Wind        struct {
			Speed float64 `json:"speed"`
			Dir   string  `json:"dir"`
			Angle int     `json:"angle"`
		} `json:"wind"`
		CloudCover struct {
			Total int `json:"total"`
		} `json:"cloud_cover"`
		Precipitation struct {
			Total float64 `json:"total"`
			Type  string  `json:"type"`
		} `json:"precipitation"`
	} `json:"data"`
}

// GetWeather ...
func GetWeather(apiKey string) (WeatherResp, error) {
	var weather WeatherResp

	requestURL := "https://www.meteosource.com/api/v1/free/point?place_id=kingsland&sections=current%2Chourly&language=en&units=us"
	req, _ := http.NewRequest("GET", requestURL, nil)

	req.Header.Add("X-API-KEY", apiKey)

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Errorf("error making request:\nStatus: %s\nError: %w\n", res.Status, err)
	}

	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&weather); err != nil {
		fmt.Printf("error decoding response body\nError: %w\n", err)
	}

	return weather, err
}

// GetWeather Gets the Meteosource API response data and transposes it to a WeatherResp struct
func GetWeatherCurrent(apiKey string) (WeatherRespCurrent, error) {
	var weather WeatherResp

	requestURL := "https://www.meteosource.com/api/v1/free/point?place_id=kingsland&sections=current%2Chourly&language=en&units=us"
	req, _ := http.NewRequest("GET", requestURL, nil)

	req.Header.Add("X-API-KEY", apiKey)

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Errorf("error making request:\nStatus: %s\nError: %w\n", res.Status, err)
	}

	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&weather); err != nil {
		fmt.Printf("error decoding response body\nError: %w\n", err)
	}

	return weather.Current, err
}

// GetWeatherHourl ...
func GetWeatherHourly(apiKey string) (WeatherRespHourly, error) {
	var weather WeatherResp

	requestURL := "https://www.meteosource.com/api/v1/free/point?place_id=kingsland&sections=current%2Chourly&language=en&units=us"
	req, _ := http.NewRequest("GET", requestURL, nil)

	req.Header.Add("X-API-KEY", apiKey)

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Errorf("Error making request:\nStatus: %s\nError: %w\n", res.Status, err)
	}

	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&weather); err != nil {
		fmt.Printf("Error decoding response body\nError: %w\n", err)

	}

	return weather.Hourly, err
}

// PrintWeather Pretty prints important Meteosource API response data
func PrintWeather(weather WeatherResp) {
	fmt.Printf("Location: Kingsland, GA\n\n")

	fmt.Printf("\tCurrent:\n")
	fmt.Printf("\t\tSummary: %s\n", weather.Current.Summary)
	fmt.Printf("\t\tTemperature: %.2f\n", weather.Current.Temperature)
	fmt.Printf("\t\tWind: %.2f mph %s", weather.Current.Wind.Speed, weather.Current.Wind.Dir)

	fmt.Printf("\tHourly:\n")
	for _, hour := range weather.Hourly.Data {
		fmt.Printf("\t\t%s\n", hour.Date)
		fmt.Printf("\t\tSummary: %s\n", hour.Summary)
		fmt.Printf("\t\tTemperature: %.2f F\n", hour.Temperature)
		fmt.Printf("\t\tPrecipitation: %s\n", hour.Precipitation.Type)
		fmt.Printf("\t\tWind: %.2f mph %s\n\n", hour.Wind.Speed, hour.Wind.Dir)
	}
}

// PrintCurrentWeather Preety prints important current weather from the Meteosource API response data
func PrintCurrentWeather(weather WeatherResp) {
	fmt.Printf("\tCurrent:\n")
	fmt.Printf("\t\tSummary: %s\n", weather.Current.Summary)
	fmt.Printf("\t\tTemperature: %.2f\n", weather.Current.Temperature)
	fmt.Printf("\t\tWind: %.2f mph %s", weather.Current.Wind.Speed, weather.Current.Wind.Dir)
}

// PrintHourlyWeather Pretty prints important hourly weather from the Meteosource API response data
func PrintHourlyWeather(weather WeatherResp) {
	fmt.Printf("\tHourly:\n")
	for _, hour := range weather.Hourly.Data {
		fmt.Printf("\t\t%s\n", hour.Date)
		fmt.Printf("\t\tSummary: %s\n", hour.Summary)
		fmt.Printf("\t\tTemperature: %.2f F\n", hour.Temperature)
		fmt.Printf("\t\tPrecipitation: %s\n", hour.Precipitation.Type)
		fmt.Printf("\t\tWind: %.2f mph %s\n\n", hour.Wind.Speed, hour.Wind.Dir)
	}
}
