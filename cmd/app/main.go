package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/ptdewey/frc-scouting-database-v2/internal/predictions"
	"github.com/ptdewey/frc-scouting-database-v2/pkg/app"
	"github.com/robfig/cron/v3"
)

// Function runAnalyzer is a helper function that is called by the cron scheduler.
// It determines which events be evaluated.
func runAnalyzer(apiKey string) {
	// Manually specify events to fetch data for
	eventKeys := []string{
		"2025vagle",
	}

	// Get event data for specified event(s)
	for _, ek := range eventKeys {
		// TODO: make the analyze event call a goroutine for parallelism
		_, err := app.AnalyzeEvent(ek, apiKey)
		if err != nil {
			fmt.Printf("Error analyzing event %s: %v\n", ek, err)
			continue

		}
	}

	// Pull statistics for currently active events
	year, err := app.AnalyzeCurrentEvents(apiKey)
	if err != nil {
		fmt.Println("Error analyzing current events.", err)
		return
	}

	// Aggregate opr data for given year
	seasonOPRs, err := app.AggregateEvents(year)
	if err != nil {
		fmt.Println("Error aggregating events.", err)
		return
	}

	// Generate predictions
	// FIX: this currently only utilizes manually specified event keys array -- need to pull all recent events (move this to multi-event?)
	for _, ek := range eventKeys {
		eventPreds, err := app.PredictCurrentEvent(seasonOPRs, ek, apiKey)
		if err != nil {
			log.Println("Failed to prediction matches from event: ", ek, err)
			continue
		}
		err = predictions.WritePredMatchesToCSV(ek+"_preds.csv", eventPreds)
		if err != nil {
			log.Println("Failed to write predictions to csv: ", ek, err)
		}
	}
}

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

	// Run analyzer on startup
	runAnalyzer(apiKey)

	// Create new cron scheduler
	c := cron.New()

	// Currently running every 2 hours
	c.AddFunc("1 8-18 * * 6,0", func() {
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
