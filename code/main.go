package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const (
	apiURL = "https://api.openweathermap.org/data/2.5/forecast"
)

func weatherHandler(w http.ResponseWriter, r *http.Request) {
	// Build the request URL
	apikey := os.Getenv("API_KEY")
	if apikey == "" {
		log.Fatal("API key not set")
	}
	url := fmt.Sprintf("%s?id=524901&appid=%s", apiURL, apikey)

	// Make the request to the weather API
	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, "Failed to fetch weather data", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Read and forward the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read weather API response", http.StatusInternalServerError)
		return
	}

	if resp.StatusCode != 200 {
		http.Error(w, fmt.Sprintf("Weather API error: %s", string(body)), resp.StatusCode)
		return
	}

	// Optionally parse and pretty-print JSON
	var parsed map[string]interface{}
	if err := json.Unmarshal(body, &parsed); err != nil {
		http.Error(w, "Failed to parse weather API response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(parsed)
}

func main() {
	http.HandleFunc("/", weatherHandler)
	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
