package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/ptdewey/frc-scouting-database-v2/pkg/app"
)


// TODO: docs
func main() {
    // Load .env file and fetch TBA api key
    err := godotenv.Load("config/.env")
    if err != nil {
        return 
    }
    apiKey := os.Getenv("API_KEY")

    // Define events to fetch data from
    eventKeys := []string {
        "2024vaash",
    }
    // TODO: loop over all current events (use EventList and dates)
    
    // Get event data for specified event(s)
    // TODO: parallelize this?
    for _, ek := range eventKeys {
        _, err := app.AnalyzeEvent(ek, apiKey)
        if err != nil {
            fmt.Printf("Error analyzing event %s: %v\n", ek, err)
            continue
        }
    }
}
