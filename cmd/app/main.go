package main

import (
	"fmt"
	"os"
    "os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/ptdewey/frc-scouting-database-v2/pkg/app"
	"github.com/robfig/cron/v3"
)


// Function main is the driver function for the statistics aggregation
// portion of the application. It determines which events data should
// be fetched for, and the schedule for when it should fetch the data.
func main() {
    // Load .env file and fetch TBA api key
    err := godotenv.Load("config/app.env")
    if err != nil {
        return 
    }
    apiKey := os.Getenv("API_KEY")

    // Create new cron scheduler
    c := cron.New()

    // Currently running every 5 minutes
    c.AddFunc("*/5 6-22 * * *", func(){
        fmt.Println("Running scheduled job...")
        runAnalyzer(apiKey)
    })

    // Start cron scheduler
    c.Start()

    // Wait until term signal is received
    sc := make(chan os.Signal, 1)
    signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM)
    <-sc

    // stop cron scheduler
    c.Stop()
}


// Function runAnalyzer is a helper function that is called by the cron scheduler.
// It determines which events be evaluated.
func runAnalyzer(apiKey string) {
    // Manually specify events to fetch data for
    eventKeys := []string {
        // "2024vagle",
    }

    // Pull statistics for currently active events
    err := app.AnalyzeCurrentEvents(apiKey)
    if err != nil {
        return 
    }

    // Get event data for specified event(s)
    for _, ek := range eventKeys {
        _, err := app.AnalyzeEvent(ek, apiKey)
        if err != nil {
            fmt.Printf("Error analyzing event %s: %v\n", ek, err)
            continue
        }
    }

}
