package analysis

import (
	"fmt"
	"testing"
)

// TestCalcEventOPR tests for correct output from CalcEventOPR.
func TestCalcEventOPR(t *testing.T) {
	// Extract alliances and scores for each match, for calculating opr and summary stats
	a, teams, err := GetEventAllianceScores(matches)
	if err != nil {
		t.Fatalf("Error fetching team alliances.")
	}

	out, err := CalcEventOPR(matches, a, teams)
	if err != nil {
		t.Fatalf("%v", err)
	}
	fmt.Println(out)
	// TODO: test output for correctness (expect some floating point variance)
}

// TestOPRToCSVRow tests for correct behavior from OPRToCSVRow
func TestOPRToCSVRow(t *testing.T) {
	return
}

// TestOPRToCSV tests for correct behavior from OPRToCSV
func TestOPRToCSV(t *testing.T) {
	return
}
