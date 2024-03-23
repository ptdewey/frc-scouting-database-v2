package aggregation

import (
	"testing"

	"github.com/ptdewey/frc-scouting-database-v2/internal/analysis"
)

// Initialize variables for testing
var(
    o analysis.OPR
    oprs []analysis.OPR
    s SeasonOPR
    omap map[string]*SeasonOPR
)

// Set up for tests
func init() {
    o = analysis.OPR{
        TeamKey: "frc2106",
        OPR: 21.57,
        AutoOPR: 9.63,
        TeleopOPR: 12.04,
        RPOPR: 0.87,
    }
    oprs = append(oprs, o)
    oprs = append(oprs, analysis.OPR{
        TeamKey: "frc2106",
        OPR: 36.71,
        AutoOPR: 19.78,
        TeleopOPR: 16.25,
        RPOPR: 1.55,
    })
    s = SeasonOPR{
        TeamKey: "frc2106",
        OPRs: []float64{ 21.57, 36.71 },
        AutoOPRs: []float64{ 9.63, 19.78 },
        TeleopOPRs: []float64{ 12.04, 16.25 },
        RPOPRs: []float64{ 0.87, 1.55 },
    } 
    omap = CombineEventOPRs(oprs)
}

// Test multievent functions for correct output
func TestMultiEvent(t *testing.T) {
    // test results from CombineEventOPRs
    team := omap["frc2106"]
    if team.TeamKey != s.TeamKey {
        t.Fatalf("Error: Unexpected result")
    }

    for i := range team.OPRs {
        if team.OPRs[i] != s.OPRs[i] {
            t.Fatalf("Error: Mismatched OPR at index %d", i)
        }
        if team.AutoOPRs[i] != s.AutoOPRs[i] {
            t.Fatalf("Error: Mismatched AutoOPR at index %d", i)
        }
        if team.TeleopOPRs[i] != s.TeleopOPRs[i] {
            t.Fatalf("Error: Mismatched TeleopOPR at index %d", i)
        }
        if team.RPOPRs[i] != s.RPOPRs[i] {
            t.Fatalf("Error: Mismatched RPOPR at index %d", i)
        }
    }

    // test calculateMaxEvents
    n := calcMaxEvents(omap)
    if n != 2 {
        t.Fatalf("Error: Mismatched max event values: %d %d", n, 2)
    }

    // test SeasonOPRtoCSV
    // TODO: get this working (pathing issues when not run from project root)
    // err := SeasonOPRtoCSV(omap, "2024")
    // if err != nil {
    //     t.Fatalf("Error: %v", err)
    // }
    
}
