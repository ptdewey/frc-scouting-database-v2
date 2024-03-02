package analysis

import (
    "encoding/csv"
	"fmt"
    "os"

	"github.com/ptdewey/frc-scouting-database-v2/internal/api"
)


// Helper function that writes an array of Matches to a csv file.
// Takes in slice of Match types and a filename for the output csv file.
// Returns the output 2D array and an error.
// This function does not modify the global state and has the side effect
// of creating a file under filename.
func MatchesToCSV(matches []api.Match, filename string) ([][]string, error) {
    // create and open new file
    f, err := os.Create(filename)
    if err != nil {
        return nil, err
    }
    defer f.Close()
    
    // create new csv file writer
    w := csv.NewWriter(f)
    defer w.Flush()
    
    // create 2D array to store rows
    var out [][]string

    // csv header
    h := []string {
        "key", "match_num", "comp_level",
        "b1", "b2", "b3", "r1", "r2", "r3",
        "b_score", "r_score",
        "b_auto_score", "r_auto_score",
        "b_tele_score", "r_tele_score",
        "b_adj_score", "r_adj_score",
        "b_rp", "r_rp", "winning_alliance",
    }
    out = append(out, h)

    // iterate through matches, adding to out
    for _, m := range matches {
        // skip loop iteration for unplayed matches with an undefined breakdown
        if m.ScoreBreakdown == nil {
            continue
        } else if m.ScoreBreakdown.Blue.TotalPoints == -1 {
            continue
        }

        out = append(out, MatchToCSVRow(m))
    }

    // write out to csv file
    if err := w.WriteAll(out); err != nil {
        return nil, err
    }
    
    return out, nil
}


// Function that converts a match type object to a string slice
// to allow writing it to a csv file.
// Takes in a match type object and outputs a string slice.
func MatchToCSVRow(m api.Match) []string {
    var out []string 

    // Output format: 
    // key, num, level, b1, b2, b3, r1, r2, r3, bscore, bauto, 
    // btele, badj, rscore, rauto, rtele, radj, brp, rrp, winner

    // Append base level match fields
    out = append(out,
        m.Key,
        fmt.Sprint(m.MatchNumber),
        m.CompLevel,
    )
    
    // Append alliance members
    out = append(out, m.Alliances.Blue.TeamKeys[0:3]...)
    out = append(out, m.Alliances.Red.TeamKeys[0:3]...)

    // Append score breakdown fields
    out = append(out, 
        fmt.Sprint(m.ScoreBreakdown.Blue.TotalPoints),
        fmt.Sprint(m.ScoreBreakdown.Red.TotalPoints),
        fmt.Sprint(m.ScoreBreakdown.Blue.AutoPoints),
        fmt.Sprint(m.ScoreBreakdown.Red.AutoPoints),
        fmt.Sprint(m.ScoreBreakdown.Blue.TeleopPoints),
        fmt.Sprint(m.ScoreBreakdown.Red.TeleopPoints),
        fmt.Sprint(m.ScoreBreakdown.Blue.AdjustPoints),
        fmt.Sprint(m.ScoreBreakdown.Red.AdjustPoints),
        fmt.Sprint(m.ScoreBreakdown.Blue.RP),
        fmt.Sprint(m.ScoreBreakdown.Red.RP),
    )

    // Append winning alliance
    out = append(out, m.WinningAlliance)

    return out
}
