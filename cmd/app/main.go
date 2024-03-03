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

    // Define events to fetch data from
    // TODO: define as all current events (use EventList and dates)
    eventKeys := []string {
        "2024vaash", "2024vabla",
        "2024caph", "2024miket", "2024paca", "2024flwp", "2024casj",
    }

    // Create new cron scheduler
    c := cron.New()

    // Currently running every 5 minutes
    c.AddFunc("*/5 8-22 * * 5,6,0", func(){
        fmt.Println("Running scheduled job...")

        // Get event data for specified event(s)
        // TODO: parallelize this?
        for _, ek := range eventKeys {
            _, err := app.AnalyzeEvent(ek, apiKey)
            if err != nil {
                fmt.Printf("Error analyzing event %s: %v\n", ek, err)
                continue
            }
        }
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
