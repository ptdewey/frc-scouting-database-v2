package analysis

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/ptdewey/frc-scouting-database-v2/internal/api"
)

// Define global variables for testing space.
var (
    apiKey string
    matches []api.Match
)

// Function init initializes variables for tests.
func init() {
    err := godotenv.Load("../../config/.env")
    if err != nil {
        fmt.Println("Error reading environment file:", err)
        return
    }
    apiKey = os.Getenv("API_KEY")

    // fetch event data
    eventKey := "2022chcmp"
    rawMatches, err := api.EventMatchesList(eventKey, apiKey) 
    if err != nil {
        fmt.Printf("Error fetching event matches. %s", eventKey)
        return
    }

    // format matches
    matches, err = api.FormatMatchList(rawMatches)
    if err != nil {
        fmt.Printf("Error formatting event matches. %s", eventKey)
        return
    }
}
