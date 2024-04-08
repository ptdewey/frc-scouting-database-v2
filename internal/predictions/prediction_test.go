package predictions

import (
	"fmt"
	"testing"

	"github.com/ptdewey/frc-scouting-database-v2/internal/aggregation"
)


// instantiate necessary variables for testing
var (
    season map[string]*aggregation.SeasonOPR
    btk = []string{ "frc2106", "frc3136", "frc404" }
    rtk = []string{ "frc346", "frc5804", "frc4821" }

    outputPath = "../../output"
    year = "2024"
)


// initialize variables for testing
func init() {
    var err error
    season, err = aggregation.ReadAndCombineEventOPRs(outputPath, year)
    if err != nil {
        fmt.Println("Error setting up OPR map in tests initialization")
        return
    }
}


// Test the PredictMatchResult function for correct behavior
func TestPredictMatchResults(t *testing.T) {
    p, err := PredictMatchResult(btk, rtk, season) 
    if err != nil {
        t.Fatalf("Error predicting match results:")
    }
    fmt.Println(p.WinningAlliance)
    fmt.Println(p.WinningMargin)
    fmt.Println(p.RedScores.Score)
    fmt.Println(p.BlueScores.Score)
}


// Test the TestPredictAllianceScore function for correct output format
func TestPredictAllianceScore(t *testing.T) {
    p, err := PredictAllianceScore(btk, season)
    if err != nil {
        t.Fatalf("Error predicting match score for alliance: %v", err)
    }

    // check output
    if len(p.TeamKeys) == 0 {
        t.Fatalf("Error: Unexpected team key length.")
    }
    if p.Score == 0 {
        t.Fatalf("Error calculating predicted score.")
    }
    if p.AutoScore == 0 {
        t.Fatalf("Error calculating predicted autonomous score.")
    }
    if p.TeleopScore == 0 {
        t.Fatalf("Error calculating predicted teleop score.")
    }
    if p.RP == 0 {
        t.Fatalf("Error calculating predicted RP score.")
    }
}


// Test getAllianceOPR for correct output format.
// Output subject to change over season so it is not directly testing for matches.
func TestAllianceOPR(t *testing.T) {
    b, err := getAllianceOPR(btk, season)
    if err != nil {
        t.Fatalf("Error extracting alliance OPR for season data: %v", err)
    }

    // test resulting alliance object
    for i := range len(b.TeamKeys) {
        if b.TeamKeys[i] != btk[i] {
            t.Fatalf("Error: Mismatched team key found at index %d.", i)
        }
        // TODO: test opr values? (they change over time so might be hard)
        // - could use a previous year's data?
    }
}


// Test getTeamOPR for incorrect behavior.
func TestGetTeamOPR(t *testing.T) {
    _, err := getTeamOPR(btk[0], season) 
    if err != nil {
        t.Fatalf("Error getting team OPR: %v", err)
    }
}
