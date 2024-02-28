package utils

import (
    "os"
	"testing"
)


// TestTeamMatches tests for correct output from TeamMatches.
// Requires working api requests and formatting systems.
func TestTeamMatches(t *testing.T) {
    tm, err := TeamMatches(&matches, "frc2106")
    if err != nil {
        t.Fatalf("%f", err)
    }

    // list of match keys for 2023 vagle frc2106 matches
    matchKeys := []string{
        "2023vagle_qm11", "2023vagle_qm19", "2023vagle_qm24", "2023vagle_qm28",
        "2023vagle_qm3", "2023vagle_qm35", "2023vagle_qm41", "2023vagle_qm52",
        "2023vagle_qm57", "2023vagle_qm63", "2023vagle_qm67", "2023vagle_qm74",
        "2023vagle_sf10m1", "2023vagle_sf12m1", "2023vagle_sf3m1", "2023vagle_sf8m1",
    }
    
    // Check match keys match correctly
    for i, m := range tm {
        if matchKeys[i] != m.Key {
            t.Fatalf("Mismatched match keys at index %d. %s %s", 
                i, matchKeys[i], m.Key)
        }
    }
}


// TestTeamMatchesToCSV tests for correct output and if csv file is correctly created.
func TestTeamMatchesToCSV(t *testing.T) {
    teamKey := "frc2106"
    filename := "test_teamutils.csv"

    // check file does not exist before writing
    _, err := os.ReadFile(filename)
    if err == nil {
        t.Fatalf("Error: file already exists. %s", filename)
    }

    // write team matches 2D array to csv file
    tm, err := TeamMatchesToCSV(&matches, teamKey, filename)
    if err != nil || tm == nil {
        t.Fatalf("%v", err)
    }

    // clean up after testing by removing file
    err = os.Remove(filename)
    if err != nil {
        t.Fatalf("Error deleting file. %s", filename)
    }
}
