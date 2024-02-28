package utils

import (
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/ptdewey/frc-scouting-database-v2/internal/api"
)


// Define global variables for testing space.
var (
    apiKey string
    matches []api.Match
    m api.Match
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
    eventKey := "2023vagle"
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

    // arbitrary match to check
    m = matches[12]
}


// Test function that converts Match type to csv row (string array)
// and test for correct output.
func TestMatchesToCSVRow(t *testing.T) {
    str := []string{
        "2023vagle_qm19", "19", "qm", "frc5724", "frc611", "frc4286",
        "frc449", "frc2998", "frc2106", "96", "106", 
        "15", "30", "61", "61", "0", "0", "0", "3", "red",
    }

    r := MatchToCSVRow(m)

    // check lengths match
    if len(r) != len(str) {
        t.Fatalf("Mismatched result lengths. %d %d", len(str), len(r))
    }

    // check all results match
    for i, s := range str {
        if s != r[i] {
            t.Fatalf("Mismatched result at index%d. %s %s", i, s, r[i])
        }
    }
}


// Test function that writes []Match type to csv file and 2D string array.
func TestMatchesToCSV(t *testing.T) {
    filename := "test_matchutils.csv"

    // check file does not exist before writing
    _, err := os.ReadFile(filename)
    if err == nil {
        t.Fatalf("Error: file already exists. %s", filename)
    }

    // write matches to csv file
    out, err := MatchesToCSV(matches, filename)
    if err != nil || out == nil {
        t.Fatalf("Error writing matches to csv file. %s", filename)
    }

    // clean up after testing by removing file
    err = os.Remove(filename)
    if err != nil {
        t.Fatalf("Error deleting file. %s", filename)
    }
}

