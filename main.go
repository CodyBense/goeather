package main

import (
	"log"
	"net/http"
	// "os"
	// "github.com/joho/godotenv"
)

type apiConfig struct {
	// apiKey string
}

func main() {
	const filePathRoot = "."
	const port = "8000"

	// err := godotenv.Load("/home/cody/workspaces/github/CodyBense/goeather/.env")
	// if err != nil {
	// 	log.Fatalf("Error loading .env file")
	// }

	// meteoKey := os.Getenv("METEOSOURCE_API_KEY")
	// if meteoKey == "" {
	// 	log.Fatalf("METEOSOURCE_API_KEY must be set")
	// }

	// apiCfg := apiConfig{
	// 	apiKey: meteoKey,
	// }
	apiCfg := apiConfig{}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/weather", apiCfg.handlerWeatherGet)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Serving on port: %s\n", port)
	log.Fatal(srv.ListenAndServe())
}
