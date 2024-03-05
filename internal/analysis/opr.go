package analysis

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"

	"github.com/ptdewey/frc-scouting-database-v2/internal/api"
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
func CalcEventOPR(matches []api.Match, a AllianceInfo, teams []string) ([]OPR, error) {
    // Create a map for quick lookup of team index based on team key.
    teamIndex := make(map[string]int)
    for i, team := range teams {
        teamIndex[team] = i
    }

    // Set number of rows and columns in matrix.
    r, c := len(a.Scores), len(teams)

    // Create data array for matrix
    data := make([]float64, r * c)

    // Fill design matrix with 1s indicating team presence in a match
    for i, a := range a.Alliances {
        for _, t := range a {
            if idx, exists := teamIndex[t]; exists {
                data[i * c + idx] = 1
            }
        }
    }

    // Create matrix, response vector, and output vector
    A := mat.NewDense(r, c, data)

    // Find expected contributions for each metric
    xs, err := oprHelper(A, a.Scores)
    if err != nil {
        return nil, err
    }
    xa, err := oprHelper(A, a.AutoScores)
    if err != nil {
        return nil, err
    }
    xt, err := oprHelper(A, a.TeleScores)
    if err != nil {
        return nil, err
    }
    xr, err := oprHelper(A, a.RP)
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

    // Sort output
    sort.Slice(oprs, func(i, j int) bool {
        return oprs[i].OPR > oprs[j].OPR
    })

    return oprs, nil
}


// Function oprHelper solves the linear system for input matrix A and
// input slice (vector) s, outputting a slice.
// Can fail in cases where the matrix is singular or nearly singular.
// It does not modify the global state and has no side effects.
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


// Function OPRToCSV converts an OPR type slice into a 2D string array
// and then writes it to a csv file called 'filename'.
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


// Function OPRToCSVRow is a helper function for OPRToCSV that converts
// a singular OPR object 'o' to a string slice, formatted as a csv file row.
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
