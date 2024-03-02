package api

import (
    "fmt"
    "os"
    "testing"
    // "io"

    "github.com/joho/godotenv"
)

// Global scope variables required for testing
var apiKey string


// Function init initializes variables for tests
func init() {
    err := godotenv.Load("../../config/app.env")
    if err != nil {
        fmt.Println("Error reading environment file:", err)
        return
    }
    apiKey = os.Getenv("API_KEY")
}


// TestEventList calls api.EventList checking for successful execution.
func TestEventList(t *testing.T) {
    year := "2023"
    out, err := EventList(year, apiKey)
    if err != nil {
        t.Fatalf(`EventList(year, apiKey) = %q. Error: %v`, out, err)
    }
}


// TestEventMatchesList calls api.EventMatchesList checking for successful execution.
func TestEventMatchesList(t *testing.T) {
    eventKey := "2024vagle"
    out, err := EventMatchesList(eventKey, apiKey)
    if err != nil {
        t.Fatalf(`EventMatches(eventKey, apiKey) = %q. Error: %v`, out, err)
    }
}


// TestTeamList calls api.TeamList checking for successful execution.
func TestTeamList(t *testing.T) {
    eventKey := "2020vagle"
    out, err := TeamList(eventKey, apiKey)
    if err != nil {
        t.Fatalf(`TeamList(eventKey, apiKey) = %q. Error: %v`, out, err)
    }
}

