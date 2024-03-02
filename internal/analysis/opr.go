package analysis

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"

	"github.com/ptdewey/frc-scouting-database-v2/internal/api"
	"github.com/ptdewey/frc-scouting-database-v2/internal/utils"
	"gonum.org/v1/gonum/mat"
)


type OPR struct {
    TeamKey string
    OPR float64
    AutoOPR float64
    TeleopOPR float64
    RPOPR float64
}


// Function that calculates expected contribution for each team at an event.
// Takes in []Match type array and outputs []OPR type array and error.
// Does not modify the global state and has no side effects.
func CalcEventOPR(matches []api.Match) ([]OPR, error) {
    // Initialize container variables
    var alliances [][]string
    var scores, autoScores, teleScores, rp []float64

    // Exit if 0 matches were played
    if len(matches) == 0 {
        fmt.Println("Match schedule does not exist.")
        return nil, errors.New("Match schedule not found for event. OPR cannot be calculated.")
    }
    
    // Iterate through matches populating alliances and scores.
    for _, m := range matches {
        // Skip matches with invalid score breakdown
        if m.ScoreBreakdown == nil {
            continue
        }

        // Remove unplayed matches and playoff matches
        if m.ScoreBreakdown.Blue.TotalPoints == -1 || m.CompLevel != "qm" {
            continue
        }

        // Append metrics to arrays
        alliances = append(alliances, m.Alliances.Blue.TeamKeys[:])
        alliances = append(alliances, m.Alliances.Red.TeamKeys[:])
        scores = append(scores, float64(m.ScoreBreakdown.Blue.TotalPoints))
        scores = append(scores, float64(m.ScoreBreakdown.Red.TotalPoints))
        autoScores = append(autoScores, float64(m.ScoreBreakdown.Blue.AutoPoints))
        autoScores = append(autoScores, float64(m.ScoreBreakdown.Red.AutoPoints))
        teleScores = append(teleScores, float64(m.ScoreBreakdown.Blue.TeleopPoints))
        teleScores = append(teleScores, float64(m.ScoreBreakdown.Red.TeleopPoints))
        rp = append(rp, float64(m.ScoreBreakdown.Blue.RP))
        rp = append(rp, float64(m.ScoreBreakdown.Red.RP))
    }

    // Extract unique team keys.
    teams := utils.ExtractUniqueStrings(alliances)

    // Create a map for quick lookup of team index based on team key.
    teamIndex := make(map[string]int)
    for i, team := range teams {
        teamIndex[team] = i
    }

    // Set number of rows and columns in matrix.
    r, c := len(scores), len(teams)

    // Create data array for matrix
    data := make([]float64, r * c)

    // Fill design matrix with 1s indicating team presence in a match
    for i, a := range alliances {
        for _, t := range a {
            if idx, exists := teamIndex[t]; exists {
                data[i * c + idx] = 1
            }
        }
    }

    // Create matrix, response vector, and output vector
    A := mat.NewDense(r, c, data)

    fmt.Println("made matrix")

    // Find expected contributions for each metric
    xs, err := oprHelper(A, scores)
    if err != nil {
        return nil, err
    }
    xa, err := oprHelper(A, autoScores)
    if err != nil {
        return nil, err
    }
    xt, err := oprHelper(A, teleScores)
    if err != nil {
        return nil, err
    }
    xr, err := oprHelper(A, rp)
    if err != nil {
        return nil, err
    }
    
    // Create and populate output struct
    oprs := make([]OPR, c)
    for i := 0; i < c; i++ {
        oprs[i].TeamKey = teams[i] 
        oprs[i].OPR = xs[i]
        oprs[i].AutoOPR = xa[i]
        oprs[i].TeleopOPR = xt[i]
        oprs[i].RPOPR = xr[i]
    }

    return oprs, nil
}


// TODO: docs
func oprHelper(A mat.Matrix, s []float64) ([]float64, error) {
    // Create response vector
    b := mat.NewVecDense(len(s), s)

    // Initialize output vector
    var x mat.VecDense

    // Solve the linear system
    if err := x.SolveVec(A, b); err != nil {
        return nil, err
    }

    // initialize oupput array
    out := make([]float64, x.Len())
    
    // Convert from vector to array
    for i := 0; i < x.Len(); i++ {
        out[i] = x.AtVec(i)
    }

    return out, nil
}



// TODO: docs
func OPRToCSV(oprs []OPR, filename string) ([][]string, error) {
    // create and open new file
    f, err := os.Create(filename)
    if err != nil {
        return nil, err
    }
    defer f.Close()

    // Create new csv file writer
    w := csv.NewWriter(f)
    defer w.Flush()

    // create 2D array to store rows
    var out [][]string
    
    // csv header
    h := []string {
        "team_key", "opr", "auto_opr", "tele_opr", "rp_opr",
    }
    out = append(out, h)

    // iterate through oprs, adding to out
    for _, m := range oprs {
        out = append(out, OPRToCSVRow(m))
    }
    
    // write out to csv file
    if err := w.WriteAll(out); err != nil {
        return nil, err
    }

    return out, nil
}


// TODO: docs
func OPRToCSVRow(o OPR) []string {
    var out []string

    // Append opr data to output row
    out = append(out,
        o.TeamKey,
        fmt.Sprint(o.OPR),
        fmt.Sprint(o.AutoOPR),
        fmt.Sprint(o.TeleopOPR),
        fmt.Sprint(o.RPOPR),
    )

    return out
}
